package crudfunctions

import (
	"databaselineservice/domain/common"
	"databaselineservice/sdk/cervello"
	"errors"
	"fmt"
)

func UseCaseLoraDeviceEntity(csvLines [][]string, keysMap map[string]int, objects []LoraDeviceInterface, parentAssetType string, entityName string, action string, token string) error {
	if common.DeleteActions[action] {
		return deleteLoraDevices(csvLines, keysMap, objects, entityName, action, token)
	}

	paginationObj := cervello.Pagination{PageSize: 9999999, PageNumber: 0}

	//region 0- get existing entities and map them
	err := objects[0].MigrateFromCsvLine(csvLines[0], keysMap)
	if err != nil {
		return err
	}
	existingDeviceMap := map[string]cervello.Device{}
	existingDevices, err := cervello.GetOrgDevicesFiltered(cervello.QueryParams{
		PaginationObj: paginationObj,
		Filters: []cervello.Filter{{
			Key:   "tags",
			Op:    "contains",
			Value: fmt.Sprintf("\"%s\"", objects[0].GetSearchTag()),
		}},
	}, token)
	if err != nil {
		return errors.New("error fetching devices")
	}
	for _, device := range existingDevices {
		existingDeviceMap[device.ID] = device
	}
	//endregion

	//region 1- migrate from csv lines
	ColumnNames := objects[0].GetEssentialKeys()
	err = common.ValidateAllColumnsExist(keysMap, ColumnNames)
	if err != nil {
		return err
	}
	newDevicesMap := map[string]LoraDeviceInterface{}
	for index, line := range csvLines {
		fmt.Printf(fmt.Sprintf("Reading Line: %d", index+2))
		err = objects[index].MigrateFromCsvLine(line, keysMap)
		if err != nil {
			fmt.Printf(fmt.Sprintf("Error reading line: %d : %s", index+2, err))
			return err
		}
		newDevicesMap[objects[index].GetGlobalId()] = objects[index]
	}
	//endregion

	//region 2- compare existing devices with new devices
	if action == common.CompareAction {
		for id := range existingDeviceMap {
			common.SleepExecution()
			if newDevicesMap[id] == nil {
				fmt.Printf(fmt.Sprintf("deleting device with id: %s", id))
				err = cervello.DeleteDevice(id, token)
				if err != nil {
					return errors.New("error deleting device")
				}
			}
		}
	}
	//endregion

	//region 3- fetch parent assets
	parentAssets, err := cervello.GetAssetsByAssetType(fmt.Sprintf("\"%s\"", parentAssetType), cervello.QueryParams{PaginationObj: paginationObj}, token)
	if err != nil {
		fmt.Printf(fmt.Sprintf("error fetching parent assets from the database: %s", err))
		return err
	}
	if len(parentAssets) <= 0 {
		fmt.Printf("no parent assets in the database")
		return errors.New("no parent assets in the database")
	}
	parentAssetMap := make(map[string]cervello.Asset)
	for _, parent := range parentAssets {
		parentAssetMap[parent.ID] = parent
	}
	//endregion

	for index := range objects {
		common.SleepExecution()

		// create device (or update if it exists)
		if fetchedDevice := existingDeviceMap[objects[index].GetGlobalId()]; fetchedDevice.ID == "" && action != common.UpdateOnlyAction {
			fmt.Printf(fmt.Sprintf("creating line no: %d", index+2))

			//region 4- set parent assets info
			parentAssetId := objects[index].GetParentAssetId()
			if parentAssetMap[parentAssetId].ID != "" {
				err = objects[index].SetParentAssetInfo(parentAssetMap[parentAssetId])
				if err != nil {
					fmt.Printf(fmt.Sprintf("error assign parent asset info for line no: %d", index+2))
					return err
				}
			} else {
				fmt.Printf(fmt.Sprintf("invalid parent asset id for line no: %d", index+2))
				return errors.New(fmt.Sprintf("invalid parent asset id for line no: %d", index+2))
			}
			//endregion

			//region 5- validate device
			if err = objects[index].Validate(); err != nil {
				fmt.Printf(fmt.Sprintf("Error validating device : %d : %s", index+2, err))
				return err
			}
			//endregion

			//region 6- create
			if err = checkUniqueLoraDeviceFeatures(objects[index]); err != nil {
				fmt.Printf(fmt.Sprintf("Error validating device : %d : %s", index+2, err))
				return err
			}

			//migrate to cervello device
			device, err := MigrateToLoraDevice(objects[index])
			if err != nil {
				fmt.Printf(fmt.Sprintf("Error migrating device : %d : %s", index+2, err))
				return err
			}

			// create device
			//createdDevice, errCreate := cervello.CreateDevice(device, token)
			createdDevice, errCreate := cervello.CreateDevice(device, token)
			if errCreate != nil {
				fmt.Printf(fmt.Sprintf("error creating line no: %d: %s", index+2, errCreate))
				return errors.New(fmt.Sprintf("error creating line no: %d: %s", index+2, errCreate))
			}

			// Activate
			err = objects[index].Activate(token)
			if err != nil {
				_ = cervello.DeleteDevice(createdDevice.ID, token)
				return errors.New("failed to activate Device: " + err.Error())
			}
			//endregion

			//region 7- assign device to application
			_, err = cervello.AssignDeviceToApplication(objects[index].GetGlobalId(), token)
			if err != nil {
				_ = cervello.DeleteDevice(objects[index].GetGlobalId(), token)
				fmt.Printf(fmt.Sprintf("error assign line no: %d to the application: %s", index+2, err))
				return errors.New(fmt.Sprintf("error assign line no: %d to the application: %s", index+2, err))
			}
			//endregion

			//region 8- assign device to parent asset
			_, err = cervello.AssignDeviceToAsset(objects[index].GetGlobalId(), objects[index].GetParentAssetId(), token)
			if err != nil {
				_ = cervello.DeleteDevice(objects[index].GetGlobalId(), token)
				fmt.Printf(fmt.Sprintf("error assign line no: %d to the parent asset: %s", index+2, err))
				return errors.New(fmt.Sprintf("error assign line no: %d to the parent asset: %s", index+2, err))
			}
			//endregion

			go common.PublishAuditLog("Create", entityName, objects[index].GetGlobalId(), objects[index])
		} else {
			if action == common.CreateOnlyAction || fetchedDevice.ID == "" {
				continue
			}
			fmt.Printf(fmt.Sprintf("updating line no: %d", index+2))

			//region 4- set parent assets info
			parentAssetId := objects[index].GetParentAssetId()
			if parentAssetMap[parentAssetId].ID != "" {
				err = objects[index].SetParentAssetInfo(parentAssetMap[parentAssetId])
				if err != nil {
					fmt.Printf(fmt.Sprintf("error assign parent asset info for line no: %d", index+2))
					return err
				}
			} else {
				fmt.Printf(fmt.Sprintf("invalid parent asset id for line no: %d", index+2))
				return errors.New(fmt.Sprintf("invalid parent asset id for line no: %d", index+2))
			}
			//endregion

			//region 5- validate device
			if err = objects[index].Validate(); err != nil {
				fmt.Printf(fmt.Sprintf("Error validating device : %d : %s", index+2, err))
				return err
			}
			//endregion

			//region 6- update device
			err = updateLoraDeviceEntity(fetchedDevice, objects[index], token)
			if err != nil {
				fmt.Printf(fmt.Sprintf("Error updating device: %d : %s", index+2, err))
				return err
			}
			//endregion

			//region 7- assign device to application
			_, err = cervello.AssignDeviceToApplication(objects[index].GetGlobalId(), token)
			if err != nil {
				_ = cervello.DeleteDevice(objects[index].GetGlobalId(), token)
				fmt.Printf(fmt.Sprintf("error assign line no: %d to the application: %s", index+2, err))
				return errors.New(fmt.Sprintf("error assign line no: %d to the application: %s", index+2, err))
			}
			//endregion

			//region 8- assign device to parent asset
			_, err = cervello.AssignDeviceToAsset(objects[index].GetGlobalId(), objects[index].GetParentAssetId(), token)
			if err != nil {
				_ = cervello.DeleteDevice(objects[index].GetGlobalId(), token)
				fmt.Printf(fmt.Sprintf("error assign line no: %d to the parent asset: %s", index+2, err))
				return errors.New(fmt.Sprintf("error assign line no: %d to the parent asset: %s", index+2, err))
			}
			//endregion

			go common.PublishAuditLog("Update", entityName, objects[index].GetGlobalId(), objects[index])
		}

	}

	return nil
}

func checkUniqueLoraDeviceFeatures(obj LoraDeviceInterface) error {
	if result, _ := common.CheckUniqueDeviceField("featureId", obj.GetFeatureId(), obj.GetLayerName()); !result {
		return errors.New("not unique featureId")
	}
	if result, _ := common.CheckUniqueDeviceField("integrationId", obj.GetIntegrationId(), ""); !result {
		return errors.New("not unique integrationId")
	}
	if result, _ := common.CheckUniqueDeviceField("mac", obj.GetMac(), ""); !result {
		return errors.New("not unique mac address")
	}
	return nil
}

func updateLoraDeviceEntity(fetchedDevice cervello.Device, obj LoraDeviceInterface, token string) error {
	// 1- check unique features
	if fetchedDevice.CustomFields["featureId"] != obj.GetFeatureId() {
		if result, _ := common.CheckUniqueDeviceField("featureId", obj.GetFeatureId(), obj.GetLayerName()); !result {
			return errors.New("not unique featureId")
		}
	}
	if fetchedDevice.CustomFields["integrationId"] != obj.GetIntegrationId() {
		if result, _ := common.CheckUniqueDeviceField("integrationId", obj.GetIntegrationId(), ""); !result {
			return errors.New("not unique integrationId")
		}
	}
	if fetchedDevice.CustomFields["mac"] != obj.GetMac() {
		if result, _ := common.CheckUniqueDeviceField("mac", obj.GetMac(), ""); !result {
			return errors.New("not unique mac address")
		}
	}

	// 2- set old activation credentials
	//obj.SetActivationInfo(fetchedDevice) // TODO: temp commented

	// 3- migrate to cervello device
	device, err := MigrateToLoraDevice(obj)
	if err != nil {
		return errors.New(fmt.Sprintf("Error migrating device : %s", err))
	}

	// 4- update device
	_ = cervello.DeleteDevice(device.ID, token)
	_, err = cervello.CreateDevice(device, token)
	if err != nil {
		return err
	}

	return nil
}

func deleteLoraDevices(csvLines [][]string, keysMap map[string]int, objects []LoraDeviceInterface, entityName string, action string, token string) error {
	paginationObj := cervello.Pagination{PageSize: 9999999, PageNumber: 0}

	//region 0- get existing entities and map them
	existingDeviceMap := map[string]cervello.Device{}

	existingDevices, err := cervello.GetOrgDevicesFiltered(cervello.QueryParams{
		PaginationObj: paginationObj,
		Filters: []cervello.Filter{{
			Key:   "tags",
			Op:    "contains",
			Value: fmt.Sprintf("\"%s\"", objects[0].GetSearchTag()),
		}},
	}, token)
	if err != nil {
		return errors.New("error fetching devices")
	}
	for _, device := range existingDevices {
		existingDeviceMap[device.ID] = device
	}
	//endregion

	//region 1- delete all objects
	if action == common.DeleteAction {
		for index, device := range existingDevices {
			fmt.Printf(fmt.Sprintf("deleting line no: %d", index+2))
			err := cervello.DeleteDevice(device.ID, token)
			if err != nil {
				return errors.New(fmt.Sprintf("error deleting line no: %d", index+2))
			}
		}
		fmt.Printf(fmt.Sprintf("all %s deleted successfully", entityName))
		return nil
	}

	//region 2- migrate from csv lines
	newDevicesMap := map[string]LoraDeviceInterface{}
	for index, line := range csvLines {
		common.SleepExecution()
		fmt.Printf(fmt.Sprintf("Reading Line: %d", index+2))
		err = objects[index].MigrateFromCsvLine(line, keysMap)
		if err != nil {
			return errors.New(fmt.Sprintf("Error reading line: %d : %s", index+2, err))
		}
		newDevicesMap[objects[index].GetGlobalId()] = objects[index]
	}
	//endregion

	//region 3- delete Other devices
	if action == common.DeleteOthersAction {
		for id := range existingDeviceMap {
			common.SleepExecution()
			if newDevicesMap[id] == nil {
				fmt.Printf(fmt.Sprintf("deleting device with id: %s", id))
				err = cervello.DeleteDevice(id, token)
				if err != nil {
					return errors.New("error deleting device")
				}
			}
		}
		fmt.Printf(fmt.Sprintf("other %s devices deleted successfully", entityName))
		return nil
	}
	//endregion

	//region 4- delete csv devices
	if action == common.DeleteCsvAction {
		for id := range newDevicesMap {
			common.SleepExecution()
			fmt.Printf(fmt.Sprintf("deleting device with id: %s", id))
			err = cervello.DeleteDevice(id, token)
			if err != nil {
				return errors.New("error deleting device")
			}

		}
		fmt.Printf(fmt.Sprintf("csv %s devices deleted successfully", entityName))
		return nil
	}
	//endregion

	return nil
}
