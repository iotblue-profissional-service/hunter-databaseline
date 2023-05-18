package infrastructuredomain

import (
	"databaselineservice/domain/common"
	"databaselineservice/domain/crudfunctions"
	"databaselineservice/sdk/cervello"
	"encoding/json"
)

func UseCaseAreas(csvLines [][]string, keysMap map[string]int, action string) (string, error) {
	interfaces := make([]crudfunctions.AssetInterface, len(csvLines))
	for index := range interfaces {
		interfaces[index] = &Area{}
	}

	if action == common.ValidateAction {
		validatingObjects := make([]crudfunctions.ValidatingObject, len(csvLines))
		for index := range interfaces {
			validatingObjects[index] = &Area{}
		}
		essentialKeys := interfaces[0].GetEssentialKeys()

		err := crudfunctions.ValidateDocument(csvLines, keysMap, essentialKeys, validatingObjects, crudfunctions.ParentsData{
			ParentAssetType:        (&City{}).GetAssetType(),
			ParentAssetKey:         interfaces[0].GetParentAssetKey(),
			ParentGatewaySearchTag: common.EmptyField,
			ParentGatewayKey:       common.EmptyField,
		}, "")
		if err != nil {
			return "validation finished", err
		}
		return "validation finished \n valid document", err
	}

	err := crudfunctions.UseCaseAssetEntity(csvLines, keysMap, interfaces, (&City{}).GetAssetType(), "area", action, "")
	if err != nil {
		return "", err
	}

	return "areas are done successfully", nil
}

func MigrateCervelloAssetToArea(asset cervello.Asset) (Area, error) {
	result := Area{}

	jsonResult, err := json.Marshal(asset.CustomFields)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(jsonResult, &result)
	if err != nil {
		return result, err
	}

	result.GlobalId = asset.ID
	result.NameArabic = asset.Name
	result.CreatedAt = asset.CreatedAt
	result.UpdatedAt = asset.UpdatedAt

	return result, err
}
