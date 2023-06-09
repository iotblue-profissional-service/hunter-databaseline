package crudfunctions

import (
	"databaselineservice/domain/common"
	"databaselineservice/sdk/cervello"
	"errors"
	"fmt"
)

func UseCaseModbusDeviceEntity(csvLines [][]string, keysMap map[string]int, objects []ModbusInterface, parentAssetType string, parentGatewayTag string, entityName string, action string, token string) error {
	if common.DeleteActions[action] {
		return deleteModbusDevices(csvLines, keysMap, objects, entityName, action, token)
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
	newDevicesMap := map[string]ModbusInterface{}
	for index, line := range csvLines {
		fmt.Println(fmt.Sprintf("Reading Line: %d", index+2))
		err = objects[index].MigrateFromCsvLine(line, keysMap)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error reading line: %d : %s", index+2, err))
			return err
		}
		newDevicesMap[objects[index].GetGlobalId()] = objects[index]
	}
	//endregion

	//region 2- fetch parent assets
	parentAssets, err := cervello.GetAssetsByAssetType(fmt.Sprintf("\"%s\"", parentAssetType), cervello.QueryParams{PaginationObj: paginationObj}, token)
	if err != nil {
		fmt.Println(fmt.Sprintf("error fetching parent assets from the database: %s", err))
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

	//region 3- compare existing devices with new devices
	if action == common.CompareAction {
		common.SleepExecution()
		for id := range existingDeviceMap {
			if newDevicesMap[id] == nil {
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
	parentGateways := make([]cervello.Device, 0)
	parentGatewayMap := make(map[string]cervello.Device)
	if parentGatewayTag != common.EmptyField {
		parentGateways, err = cervello.GetOrgDevicesFiltered(cervello.QueryParams{
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

		// create device (or update if it exists)
		if fetchedDevice := existingDeviceMap[objects[index].GetGlobalId()]; fetchedDevice.ID == "" && action != common.UpdateOnlyAction {
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
			if err = checkUniqueModbusFeatures(objects[index]); err != nil {
				fmt.Println(fmt.Sprintf("Error validating device : %d : %s", index+2, err))
				return err
			}

			//migrate to cervello device
			device, err := MigrateToCervelloModBusDevice(objects[index])
			if err != nil {
				fmt.Println(fmt.Sprintf("Error migrating device : %d : %s", index+2, err))
				return err
			}
			if objects[index].GetDeviceType() == cervello.DeviceTypeGateWay {
				device.CommunicationProtocol = cervello.DeviceProtocolModbus
				// device.ProtocolConfigurations = cervello.ModbusDeviceConfig{}
			}

			// create device
			_, err = cervello.CreateDevice(device, token)
			if err != nil {
				fmt.Println(fmt.Sprintf("error creating line no: %d: %s", index+2, err))
				return errors.New(fmt.Sprintf("error creating line no: %d: %s", index+2, err))
			}
			//endregion

			//region 9- assign device to application
			_, err = cervello.AssignDeviceToApplication(objects[index].GetGlobalId(), token)
			if err != nil {
				_ = cervello.DeleteDevice(objects[index].GetGlobalId(), token)
				fmt.Println(fmt.Sprintf("error assign line no: %d to the application: %s", index+2, err))
				return errors.New(fmt.Sprintf("error assign line no: %d to the application: %s", index+2, err))
			}
			//endregion

			//region 10- assign device to parent asset
			_, err = cervello.AssignDeviceToAsset(objects[index].GetGlobalId(), objects[index].GetParentAssetId(), token)
			if err != nil {
				_ = cervello.DeleteDevice(objects[index].GetGlobalId(), token)
				fmt.Println(fmt.Sprintf("error assign line no: %d to the parent asset: %s", index+2, err))
				return errors.New(fmt.Sprintf("error assign line no: %d to the parent asset: %s", index+2, err))
			}
			//endregion

			//region 11- set modbus config
			if objects[index].GetDeviceType() == cervello.DeviceTypeGateWay {
				err = cervello.UpdateModbusDeviceConfig(objects[index].GetGlobalId(), objects[index].GetModbusConfig(), "")
				if err != nil {
					_ = cervello.DeleteDevice(objects[index].GetGlobalId(), token)
					fmt.Println(fmt.Sprintf("error set modbus config for line no: %d to the parent asset: %s", index+2, err))
					return errors.New(fmt.Sprintf("error set modbus config for line no: %d to the parent asset: %s", index+2, err))
				}
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
			err = updateModbusEntity(fetchedDevice, objects[index], token)
			if err != nil {
				fmt.Println(fmt.Sprintf("Error updating device: %d : %s", index+2, err))
				return err
			}
			//endregion

			//region 9- assign device to application
			_, err = cervello.AssignDeviceToApplication(objects[index].GetGlobalId(), token)
			if err != nil {
				_ = cervello.DeleteDevice(objects[index].GetGlobalId(), token)
				fmt.Println(fmt.Sprintf("error assign line no: %d to the application: %s", index+2, err))
				return errors.New(fmt.Sprintf("error assign line no: %d to the application: %s", index+2, err))
			}
			//endregion

			//region 10- assign device to parent asset
			_, err = cervello.AssignDeviceToAsset(objects[index].GetGlobalId(), objects[index].GetParentAssetId(), token)
			if err != nil {
				_ = cervello.DeleteDevice(objects[index].GetGlobalId(), token)
				fmt.Println(fmt.Sprintf("error assign line no: %d to the parent asset: %s", index+2, err))
				return errors.New(fmt.Sprintf("error assign line no: %d to the parent asset: %s", index+2, err))
			}
			//endregion

			//region 11- set modbus config
			if objects[index].GetDeviceType() == cervello.DeviceTypeGateWay {
				err = cervello.UpdateModbusDeviceConfig(objects[index].GetGlobalId(), objects[index].GetModbusConfig(), "")
				if err != nil {
					_ = cervello.DeleteDevice(objects[index].GetGlobalId(), token)
					fmt.Println(fmt.Sprintf("error set modbus config for line no: %d to the parent asset: %s", index+2, err))
					return errors.New(fmt.Sprintf("error set modbus config for line no: %d to the parent asset: %s", index+2, err))
				}
			}
			//endregion

			go common.PublishAuditLog("Update", entityName, objects[index].GetGlobalId(), objects[index])
		}

	}

	return nil
}

func checkUniqueModbusFeatures(obj ModbusInterface) error {
	//if result, _ := common.CheckUniqueDeviceField("featureId", obj.GetFeatureId(), obj.GetLayerName()); !result {
	//	return errors.New("not unique featureId")
	//}
	if result, _ := common.CheckUniqueDeviceField("integrationId", obj.GetReferenceName(), ""); !result {
		return errors.New("not unique integrationId")
	}
	if obj.GetIP() != common.EmptyField {
		if result, _ := common.CheckUniqueDeviceField("ip", obj.GetIP(), ""); !result {
			return errors.New("not unique ip")
		}
	}
	if obj.GetMac() != common.EmptyField {
		if result, _ := common.CheckUniqueDeviceField("mac", obj.GetMac(), ""); !result {
			return errors.New("not unique mac")
		}
	}
	return nil
}

func updateModbusEntity(fetchedDevice cervello.Device, obj ModbusInterface, token string) error {
	// 1- check unique features
	//if fetchedDevice.CustomFields["featureId"] != obj.GetFeatureId() {
	//	if result, _ := common.CheckUniqueDeviceField("featureId", obj.GetFeatureId(), obj.GetLayerName()); !result {
	//		return errors.New("not unique featureId")
	//	}
	//}
	if fetchedDevice.CustomFields["integrationId"] != obj.GetReferenceName() {
		if result, _ := common.CheckUniqueDeviceField("integrationId", obj.GetReferenceName(), ""); !result {
			return errors.New("not unique integrationId")
		}
	}
	if obj.GetIP() != common.EmptyField {
		if fetchedDevice.CustomFields["ip"] != obj.GetIP() {
			if result, _ := common.CheckUniqueDeviceField("ip", obj.GetIP(), ""); !result {
				return errors.New("not unique ip")
			}
		}
	}
	if obj.GetMac() != common.EmptyField {
		if fetchedDevice.CustomFields["mac"] != obj.GetMac() {
			if result, _ := common.CheckUniqueDeviceField("mac", obj.GetMac(), ""); !result {
				return errors.New("not unique ip")
			}
		}
	}

	// 2- migrate to cervello device
	device, err := MigrateToCervelloModBusDevice(obj)
	if err != nil {
		return errors.New(fmt.Sprintf("Error migrating device : %s", err))
	}
	if obj.GetDeviceType() == cervello.DeviceTypeGateWay {
		device.CommunicationProtocol = cervello.DeviceProtocolModbus
	}

	// 3- update device
	err = cervello.UpdateDevice(device.ID, device, token)
	if err != nil {
		return err
	}
	return nil
}

func deleteModbusDevices(csvLines [][]string, keysMap map[string]int, objects []ModbusInterface, entityName string, action string, token string) error {
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
			fmt.Println(fmt.Sprintf("deleting line no: %d", index+2))
			err := cervello.DeleteDevice(device.ID, token)
			if err != nil {
				return errors.New(fmt.Sprintf("error deleting line no: %d", index+2))
			}
		}
		fmt.Println(fmt.Sprintf("all %s deleted successfully", entityName))
		return nil
	}

	//region 2- migrate from csv lines
	ColumnNames := objects[0].GetEssentialKeys()
	err = common.ValidateAllColumnsExist(keysMap, ColumnNames)
	if err != nil {
		return err
	}
	newDevicesMap := map[string]ModbusInterface{}
	for index, line := range csvLines {
		common.SleepExecution()
		fmt.Println(fmt.Sprintf("Reading Line: %d", index+2))
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
				fmt.Println(fmt.Sprintf("deleting device with id: %s", id))
				err = cervello.DeleteDevice(id, token)
				if err != nil {
					return errors.New("error deleting device")
				}
			}
		}
		fmt.Println(fmt.Sprintf("other %s devices deleted successfully", entityName))
		return nil
	}
	//endregion

	//region 4- delete csv devices
	if action == common.DeleteCsvAction {
		for id := range newDevicesMap {
			common.SleepExecution()
			fmt.Println(fmt.Sprintf("deleting device with id: %s", id))
			err = cervello.DeleteDevice(id, token)
			if err != nil {
				return errors.New("error deleting device")
			}

		}
		fmt.Println(fmt.Sprintf("csv %s devices deleted successfully", entityName))
		return nil
	}
	//endregion

	return nil
}
