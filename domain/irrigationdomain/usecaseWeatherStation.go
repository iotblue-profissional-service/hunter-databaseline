package irrigationdomain

import (
	"databaselineservice/domain/common"
	"databaselineservice/domain/crudfunctions"
	"databaselineservice/domain/infrastructuredomain"
)

func UseCaseWeatherStation(csvLines [][]string, keysMap map[string]int, action string) (string, error) {
	interfaces := make([]crudfunctions.ModbusInterface, len(csvLines))
	for index := range interfaces {
		interfaces[index] = &WeatherStation{}
	}

	if action == common.ValidateAction {
		validatingObjects := make([]crudfunctions.ValidatingObject, len(csvLines))
		for index := range interfaces {
			validatingObjects[index] = &WeatherStation{}
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

	err := crudfunctions.UseCaseModbusDeviceEntity(csvLines, keysMap, interfaces, (&infrastructuredomain.Area{}).GetAssetType(), common.EmptyField, "WeatherStation", action, "")
	if err != nil {
		return "", err
	}

	return "WeatherStations are done successfully", nil
}
