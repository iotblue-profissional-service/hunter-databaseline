package irrigationdomain

import (
	"databaselineservice/domain/common"
	"databaselineservice/domain/crudfunctions"
	"databaselineservice/domain/infrastructuredomain"
)

func UseCaseHunterFlowZone(csvLines [][]string, keysMap map[string]int, action string) (string, error) {
	interfaces := make([]crudfunctions.DeviceInterface, len(csvLines))
	for index := range interfaces {
		interfaces[index] = &HunterFlowZone{}
	}

	if action == common.ValidateAction {
		validatingObjects := make([]crudfunctions.ValidatingObject, len(csvLines))
		for index := range interfaces {
			validatingObjects[index] = &HunterFlowZone{}
		}
		essentialKeys := interfaces[0].GetEssentialKeys()

		err := crudfunctions.ValidateDocument(csvLines, keysMap, essentialKeys, validatingObjects, crudfunctions.ParentsData{
			ParentAssetType:        (&infrastructuredomain.Area{}).GetAssetType(),
			ParentAssetKey:         interfaces[0].GetParentAssetKey(),
			ParentGatewaySearchTag: (&HunterController{}).GetSearchTag(),
			ParentGatewayKey:       interfaces[0].GetParentGatewayKey(),
		}, "")
		if err != nil {
			return "validation finished", err
		}
		return "validation finished \n valid document", err
	}

	err := crudfunctions.UseCaseDeviceEntity(csvLines, keysMap, interfaces, (&infrastructuredomain.Area{}).GetAssetType(), (&HunterController{}).GetSearchTag(), "HunterFlowZone", action, "")
	if err != nil {
		return "", err
	}

	return "Hunter Flow Zones are done successfully", nil
}
