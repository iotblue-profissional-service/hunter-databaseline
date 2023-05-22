package crudfunctions

import (
	"databaselineservice/domain/common"
	"databaselineservice/sdk/cervello"
	"errors"
	"fmt"
)

func UseCaseDeviceAssetEntity(csvLines [][]string, keysMap map[string]int, objects []DeviceAssetInterface, parentAssetType string, parentGatewayTag string, entityName string, action string, token string) error {
	if common.DeleteActions[action] {
		objectDevices := make([]DeviceInterface, len(objects))
		objectAssets := make([]AssetInterface, len(objects))
		for i := range objects {
			objectDevices[i] = objects[i]
			objectAssets[i] = objects[i]
		}
		err := deleteDevices(csvLines, keysMap, objectDevices, entityName, action, token)
		if err != nil {
			return err
		}
		err = deleteAssets(csvLines, keysMap, objectAssets, entityName, action, token)
		if err != nil {
			return err
		}
		return nil
	}

	paginationObj := cervello.Pagination{PageSize: 9999999, PageNumber: 0}

	//region 0- get existing entities and map them
	err := objects[0].MigrateFromCsvLine(csvLines[0], keysMap)
	if err != nil {
		return err
	}
	existingAssetsMap := map[string]cervello.Asset{}
	existingAssets, err := cervello.GetAssetsByAssetType(fmt.Sprintf("\"%s\"", objects[0].GetAssetType()), cervello.QueryParams{
		PaginationObj: paginationObj,
	}, token)
	if err != nil {
		return errors.New("error fetching assets")
	}
	for _, asset := range existingAssets {
		existingAssetsMap[asset.ID] = asset
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
	newAssetsMap := map[string]AssetInterface{}
	newDevicesMap := map[string]DeviceInterface{}
	for index, line := range csvLines {
		fmt.Println(fmt.Sprintf("Reading Line: %d", index+2))
		err := objects[index].MigrateFromCsvLine(line, keysMap)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error reading line: %d : %s", index+2, err))
			return err
		}
		newAssetsMap[objects[index].GetGlobalId()] = objects[index]
		newDevicesMap[objects[index].GetGlobalId()] = objects[index]
	}
	//endregion

	//region 2- fetch parent assets
	parentAssets, err := cervello.GetAssetsByAssetType(fmt.Sprintf("\"%s\"", parentAssetType), cervello.QueryParams{PaginationObj: paginationObj}, token)
	if err != nil {
		fmt.Println(fmt.Sprintf("Errors fetching parent assets from the database: %s", err))
		return err
	}
	if len(parentAssets) <= 0 {
		fmt.Println("no parent assets in the database")
		return errors.New("no parent assets in the database")
	}
	parentAssetMap := make(map[string]cervello.Asset)
	for _, parent := range parentAssets {
		parentAssetMap[parent.ID] = parent
	}
	//endregion

	//region 3- compare existing devices & assets with new ones
	if action == common.CompareAction {
		for id := range existingAssetsMap {
			common.SleepExecution()
			if newAssetsMap[id] == nil {
				fmt.Println(fmt.Sprintf("deleting asset with id: %s", id))
				err = cervello.DeleteAsset(id, token)
				if err != nil {
					return errors.New("error deleting asset")
				}
				fmt.Println(fmt.Sprintf("deleting device with id: %s", id))
				err = cervello.DeleteDevice(id, token)
				if err != nil {
					return errors.New("error deleting device")
				}
			}
		}
	}
	//endregion

	//region 4- fetch parent gateways
	// parentGateways := make([]cervello.Device, 0)
	parentGatewayMap := make(map[string]cervello.Device)
	if parentGatewayTag != common.EmptyField {
		parentGateways, err := cervello.GetOrgDevicesFiltered(cervello.QueryParams{
			PaginationObj: paginationObj,
			Filters: []cervello.Filter{{
				Key:   "tags",
				Op:    "contains",
				Value: fmt.Sprintf("\"%s\"", parentGatewayTag),
			}},
		}, token)
		if (err != nil) || (len(parentGateways) == 0) {
			fmt.Println("no parent gateways in the database")
			return errors.New("no parent gateways in the database")
		}
		for _, parent := range parentGateways {
			parentGatewayMap[parent.ID] = parent
		}
	}
	//endregion

	for index := range objects {
		common.SleepExecution()

		// create device & asset (or update if it exists)
		fetchedAsset := existingAssetsMap[objects[index].GetGlobalId()]
		fetchedDevice := existingDeviceMap[objects[index].GetGlobalId()]
		if fetchedAsset.ID == "" && fetchedDevice.ID == "" && action != common.UpdateOnlyAction {
			fmt.Println(fmt.Sprintf("creating line no: %d", index+2))
			//region 5- set parent assets info
			parentAssetId := objects[index].GetParentAssetId()
			if parentAssetMap[parentAssetId].ID != "" {
				err = objects[index].SetParentAssetInfo(parentAssetMap[parentAssetId])
				if err != nil {
					fmt.Println(fmt.Sprintf("error assign parent asset info for line no: %d", index+2))
					return err
				}
			} else {
				fmt.Println(fmt.Sprintf("invalid parent asset id for line no: %d", index+2))
				return errors.New(fmt.Sprintf("invalid parent asset id for line no: %d", index+2))
			}
			//endregion

			//region 6- set parent gateways info
			if parentGatewayTag != common.EmptyField {
				parentGatewayId := objects[index].GetParentGatewayId()
				if parentGatewayMap[parentGatewayId].ID != "" {
					err = objects[index].SetParentGatewayInfo(parentGatewayMap[parentGatewayId])
					if err != nil {
						fmt.Println(fmt.Sprintf("error assign parent gateway info for line no: %d", index+2))
						return err
					}
				} else {
					fmt.Println(fmt.Sprintf("invalid parent gateway id for line no: %d", index+2))
					return errors.New(fmt.Sprintf("invalid gateway asset id for line no: %d", index+2))
				}
			}
			//endregion

			//region 7- validate device
			if err = objects[index].Validate(); err != nil {
				fmt.Println(fmt.Sprintf("Error validating device : %d : %s", index+2, err))
				return err
			}
			//endregion

			//region 8- create
			if err = checkUniqueDeviceFeatures(objects[index]); err != nil {
				fmt.Println(fmt.Sprintf("Error validating device : %d : %s", index+2, err))
				return err
			}
			if err = checkUniqueAssetFeatures(objects[index]); err != nil {
				fmt.Println(fmt.Sprintf("Error validating asset : %d : %s", index+2, err))
				return err
			}
			if objects[index].GetDeviceType() == cervello.DeviceTypeGateWay {
				if result, _ := common.CheckUniqueAssetField("ip", objects[index].GetIP(), ""); !result {
					return errors.New(fmt.Sprintf("Error validating device : %d : %s", index+2, "not unique IP"))
				}
			}

			//migrate to cervello device & asset
			device, err := MigrateToCervelloDevice(objects[index])
			if err != nil {
				fmt.Println(fmt.Sprintf("Error migrating device : %d : %s", index+2, err))
				return err
			}
			asset, err := MigrateToCervelloAsset(objects[index])
			if err != nil {
				fmt.Println(fmt.Sprintf("Error migrating asset : %d : %s", index+2, err))
				return err
			}
			// create device
			createdDevice, err := cervello.CreateDevice(device, token)
			if err != nil {
				fmt.Println(fmt.Sprintf("Error creating line no: %d: %s", index+2, err))
				return errors.New(fmt.Sprintf("error creating line no: %d: %s", index+2, err))
			}
			// create asset
			_, err = cervello.CreateAsset(asset, token)
			if err != nil {
				fmt.Println(fmt.Sprintf("error creating line no: %d: %s", index+2, err))
				_ = cervello.DeleteDevice(createdDevice.ID, token)
				return errors.New(fmt.Sprintf("error creating line no: %d: %s", index+2, err))
			}
			//endregion

			//region 9- assign device to application
			_, err = cervello.AssignDeviceToApplication(objects[index].GetGlobalId(), token)
			if err != nil {
				_ = cervello.DeleteDevice(objects[index].GetGlobalId(), token)
				_ = cervello.DeleteAsset(objects[index].GetGlobalId(), token)
				fmt.Println(fmt.Sprintf("error assign line no: %d to the application: %s", index+2, err))
				return errors.New(fmt.Sprintf("error assign line no: %d to the application: %s", index+2, err))
			}
			//endregion

			//region 10- assign device to parent asset
			_, err = cervello.AssignDeviceToAsset(objects[index].GetGlobalId(), objects[index].GetParentAssetId(), token)
			if err != nil {
				_ = cervello.DeleteDevice(objects[index].GetGlobalId(), token)
				_ = cervello.DeleteAsset(objects[index].GetGlobalId(), token)
				fmt.Println(fmt.Sprintf("error assign line no: %d to the parent asset", index))
				return errors.New(fmt.Sprintf("error assign line no: %d to the parent asset", index))
			}
			//endregion

			//region 11- assign asset to parent asset
			_, err = cervello.AssignAssetToAsset(objects[index].GetGlobalId(), objects[index].GetParentAssetId(), token)
			if err != nil {
				_ = cervello.DeleteDevice(objects[index].GetGlobalId(), token)
				_ = cervello.DeleteAsset(objects[index].GetGlobalId(), token)
				fmt.Println(fmt.Sprintf("error assign line no: %d to the parent asset: %s", index+2, err))
				return errors.New(fmt.Sprintf("error assign line no: %d to the parent asset: %s", index+2, err))
			}
			//endregion

			go common.PublishAuditLog("Create", entityName, objects[index].GetGlobalId(), objects[index])
		} else {
			if action == common.CreateOnlyAction || fetchedDevice.ID == "" {
				continue
			}
			fmt.Println(fmt.Sprintf("updating line no: %d", index+2))
			//region 5- set parent assets info
			parentAssetId := objects[index].GetParentAssetId()
			if parentAssetMap[parentAssetId].ID != "" {
				err = objects[index].SetParentAssetInfo(parentAssetMap[parentAssetId])
				if err != nil {
					fmt.Println(fmt.Sprintf("error assign parent asset info for line no: %d", index+2))
					return err
				}
			} else {
				fmt.Println(fmt.Sprintf("invalid parent asset id for line no: %d", index+2))
				return errors.New(fmt.Sprintf("invalid parent asset id for line no: %d", index+2))
			}
			//endregion

			//region 6- set parent gateways info
			if parentGatewayTag != common.EmptyField {
				parentGatewayId := objects[index].GetParentGatewayId()
				if parentGatewayMap[parentGatewayId].ID != "" {
					err = objects[index].SetParentGatewayInfo(parentGatewayMap[parentGatewayId])
					if err != nil {
						fmt.Println(fmt.Sprintf("error assign parent gateway info for line no: %d", index+2))
						return err
					}
				} else {
					fmt.Println(fmt.Sprintf("invalid parent gateway id for line no: %d", index+2))
					return errors.New(fmt.Sprintf("invalid gateway asset id for line no: %d", index+2))
				}
			}
			//endregion

			//region 7- validate device
			if err = objects[index].Validate(); err != nil {
				fmt.Println(fmt.Sprintf("Error validating device : %d : %s", index+2, err))
				return err
			}
			//endregion

			//region 8- update
			err = UpdateDeviceAssetEntity(fetchedAsset, objects[index], token)
			if err != nil {
				fmt.Println(fmt.Sprintf("Error updating device: %d : %s", index+2, err))
				return err
			}
			//endregion

			//region 9- assign device to application
			_, err = cervello.AssignDeviceToApplication(objects[index].GetGlobalId(), token)
			if err != nil {
				_ = cervello.DeleteDevice(objects[index].GetGlobalId(), token)
				_ = cervello.DeleteAsset(objects[index].GetGlobalId(), token)
				fmt.Println(fmt.Sprintf("error assign line no: %d to the application: %s", index+2, err))
				return errors.New(fmt.Sprintf("error assign line no: %d to the application: %s", index+2, err))
			}
			//endregion

			//region 10- assign device to parent asset
			_, err = cervello.AssignDeviceToAsset(objects[index].GetGlobalId(), objects[index].GetParentAssetId(), token)
			if err != nil {
				_ = cervello.DeleteDevice(objects[index].GetGlobalId(), token)
				_ = cervello.DeleteAsset(objects[index].GetGlobalId(), token)
				fmt.Println(fmt.Sprintf("error assign line no: %d to the parent asset", index))
				return errors.New(fmt.Sprintf("error assign line no: %d to the parent asset", index))
			}
			//endregion

			//region 11- assign asset to parent asset
			_, err = cervello.AssignAssetToAsset(objects[index].GetGlobalId(), objects[index].GetParentAssetId(), token)
			if err != nil {
				_ = cervello.DeleteDevice(objects[index].GetGlobalId(), token)
				_ = cervello.DeleteAsset(objects[index].GetGlobalId(), token)
				fmt.Println(fmt.Sprintf("error assign line no: %d to the parent asset: %s", index+2, err))
				return errors.New(fmt.Sprintf("error assign line no: %d to the parent asset: %s", index+2, err))
			}
			//endregion
			go common.PublishAuditLog("Update", entityName, objects[index].GetGlobalId(), objects[index])
		}

	}

	return nil
}

func UpdateDeviceAssetEntity(fetchedAsset cervello.Asset, obj DeviceAssetInterface, token string) error {
	// 1- check unique features
	if fetchedAsset.CustomFields["featureId"] != obj.GetFeatureId() {
		if result, _ := common.CheckUniqueDeviceField("featureId", obj.GetFeatureId(), obj.GetSearchTag()); !result {
			return errors.New("not unique featureId")
		}
	}
	if fetchedAsset.CustomFields["integrationId"] != obj.GetReferenceName() {
		if result, _ := common.CheckUniqueDeviceField("integrationId", obj.GetReferenceName(), ""); !result {
			return errors.New("not unique integrationId")
		}
	}
	if obj.GetIP() != common.EmptyField {
		if fetchedAsset.CustomFields["ip"] != obj.GetIP() {
			if result, _ := common.CheckUniqueDeviceField("ip", obj.GetIP(), ""); !result {
				return errors.New("not unique ip")
			}
		}
	}
	if obj.GetMac() != common.EmptyField {
		if fetchedAsset.CustomFields["mac"] != obj.GetMac() {
			if result, _ := common.CheckUniqueDeviceField("mac", obj.GetMac(), ""); !result {
				return errors.New("not unique ip")
			}
		}
	}
	//
	if fetchedAsset.CustomFields["featureId"] != obj.GetFeatureId() {
		if result, _ := common.CheckUniqueAssetField("featureId", obj.GetFeatureId(), obj.GetSearchTag()); !result {
			return errors.New("not unique featureId")
		}
	}
	if fetchedAsset.CustomFields["integrationId"] != obj.GetReferenceName() {
		if result, _ := common.CheckUniqueAssetField("integrationId", obj.GetReferenceName(), ""); !result {
			return errors.New("not unique integrationId")
		}
	}
	if obj.GetDeviceType() == cervello.DeviceTypeGateWay {
		if fetchedAsset.CustomFields["ip"] != obj.GetIP() {
			if result, _ := common.CheckUniqueAssetField("ip", obj.GetIP(), ""); !result {
				return errors.New("not unique ip")
			}
		}
	}

	// 2- migrate to cervello device & asset
	device, err := MigrateToCervelloDevice(obj)
	if err != nil {
		return errors.New(fmt.Sprintf("Error migrating device : %s", err))
	}
	asset, err := MigrateToCervelloAsset(obj)
	if err != nil {
		return errors.New(fmt.Sprintf("Error migrating asset : %s", err))
	}

	// 3- update device & asset
	err = cervello.UpdateAsset(asset.ID, asset, token)
	if err != nil {
		return err
	}

	err = cervello.UpdateDevice(device.ID, device, token)
	if err != nil {
		// rollback updated asset
		err := cervello.UpdateAsset(asset.ID, fetchedAsset, token)
		return err
	}

	return nil
}
