package irrigationdomain

import (
	"databaselineservice/domain/common"
	"databaselineservice/domain/crudfunctions"
	"databaselineservice/domain/infrastructuredomain"
	"databaselineservice/sdk/cervello"
	"encoding/json"
)

func UseCaseHunterController(csvLines [][]string, keysMap map[string]int, action string) (string, error) {
	interfaces := make([]crudfunctions.ModbusInterface, len(csvLines))
	for index := range interfaces {
		interfaces[index] = &HunterController{}
	}

	if action == common.ValidateAction {
		validatingObjects := make([]crudfunctions.ValidatingObject, len(csvLines))
		for index := range interfaces {
			validatingObjects[index] = &HunterController{}
		}
		essentialKeys := interfaces[0].GetEssentialKeys()

		err := crudfunctions.ValidateDocument(csvLines, keysMap, essentialKeys, validatingObjects, crudfunctions.ParentsData{
			ParentAssetType:        (&infrastructuredomain.Area{}).GetAssetType(),
			ParentAssetKey:         interfaces[0].GetParentAssetKey(),
			ParentGatewaySearchTag: common.EmptyField,
			ParentGatewayKey:       interfaces[0].GetParentGatewayKey(),
		}, "")
		if err != nil {
			return "validation finished", err
		}
		return "validation finished \n valid document", err
	}

	err := crudfunctions.UseCaseModbusDeviceEntity(csvLines, keysMap, interfaces, (&infrastructuredomain.Area{}).GetAssetType(), common.EmptyField, "HunterController", action, "")
	if err != nil {
		return "", err
	}

	return "HunterControllers are done successfully", nil
}

func MigrateHunterControllerFromCervelloDevice(device cervello.Device) (HunterController, error) {
	result := HunterController{}

	jsonResult, err := json.Marshal(device.CustomFields)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(jsonResult, &result)
	if err != nil {
		return result, err
	}

	result.GlobalId = device.ID
	result.Name = device.Name
	result.IntegrationId = device.ReferenceName
	result.CreatedAt = device.CreatedAt
	result.UpdatedAt = device.UpdatedAt

	return result, err
}
