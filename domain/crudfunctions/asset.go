package crudfunctions

import (
	"databaselineservice/domain/common"
	"databaselineservice/sdk/cervello"
	"errors"
	"fmt"
)

func UseCaseAssetEntity(csvLines [][]string, keysMap map[string]int, objects []AssetInterface, parentAssetType string, entityName string, action string, token string) error {
	if common.DeleteActions[action] {
		return deleteAssets(csvLines, keysMap, objects, entityName, action, token)
	}

	//region 0- get existing entities and map them
	existingAssetsMap := map[string]cervello.Asset{}
	existingAssets, err := cervello.GetAssetsByAssetType(fmt.Sprintf("\"%s\"", objects[0].GetAssetType()), cervello.QueryParams{
		PaginationObj: cervello.Pagination{
			PageNumber: 1,
			PageSize:   99999999999,
		},
	}, token)
	if err != nil {
		return errors.New("error fetching assets")
	}
	for _, asset := range existingAssets {
		existingAssetsMap[asset.ID] = asset
	}
	//endregion

	//region 1- migrate from csv lines & make new assets map
	ColumnNames := objects[0].GetEssentialKeys()
	err = common.ValidateAllColumnsExist(keysMap, ColumnNames)
	if err != nil {
		return err
	}
	newAssetsMap := map[string]AssetInterface{}
	for index, line := range csvLines {
		fmt.Printf(fmt.Sprintf("Reading Line: %d", index+2))
		err = objects[index].MigrateFromCsvLine(line, keysMap)
		if err != nil {
			fmt.Printf(fmt.Sprintf("Error reading line: %d : %s", index+2, err))
			return err
		}
		newAssetsMap[objects[index].GetGlobalId()] = objects[index]
	}
	//endregion

	//region 2- compare existing assets with new assets
	if action == common.CompareAction {
		for id := range existingAssetsMap {
			common.SleepExecution()
			if newAssetsMap[id] == nil {
				fmt.Printf(fmt.Sprintf("deleting asset with id: %s", id))
				err = cervello.DeleteAsset(id, token)
				if err != nil {
					return errors.New("error deleting asset")
				}
			}
		}
	}
	//endregion

	//region 3- fetch parent assets
	paginationObj := cervello.Pagination{PageSize: 999999, PageNumber: 0}
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

		//create asset (or update if it exists)
		fetchedAsset := existingAssetsMap[objects[index].GetGlobalId()]
		if fetchedAsset.ID == "" && action != common.UpdateOnlyAction {
			fmt.Printf(fmt.Sprintf("creating line no: %d", index+2))
			//region 4-set parent assets info
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

			//region 5- validate asset
			if err = objects[index].Validate(); err != nil {
				fmt.Printf(fmt.Sprintf("Error validating asset : %d : %s", index+2, err))
				return err
			}
			//endregion

			//region 6- create
			if err = checkUniqueAssetFeatures(objects[index]); err != nil {
				fmt.Printf(fmt.Sprintf("Error validating asset : %d : %s", index+2, err))
				return err
			}

			//migrate to cervello asset
			asset, err := MigrateToCervelloAsset(objects[index])
			if err != nil {
				fmt.Printf(fmt.Sprintf("Error migrating asset : %d : %s", index+2, err))
				return err
			}

			// create asset
			_, err = cervello.CreateAsset(asset, token)
			if err != nil {
				fmt.Printf(fmt.Sprintf("error creating line no: %d : %s", index+2, err))
				return errors.New(fmt.Sprintf("error creating line no: %d : %s", index+2, err))
			}
			//endregion

			//region 7- assign asset to parent asset
			_, err = cervello.AssignAssetToAsset(objects[index].GetGlobalId(), objects[index].GetParentAssetId(), token)
			if err != nil {
				_ = cervello.DeleteAsset(objects[index].GetGlobalId(), token)
				fmt.Printf(fmt.Sprintf("error assign line no: %d to the parent asset: %s", index+2, err))
				return errors.New(fmt.Sprintf("error assign line no: %d to the parent asset: %s", index+2, err))
			}
			//endregion

			go common.PublishAuditLog("Create", entityName, objects[index].GetGlobalId(), objects[index])
		} else {
			if action == common.CreateOnlyAction || fetchedAsset.ID == "" {
				continue
			}

			fmt.Printf(fmt.Sprintf("updating line no: %d", index+2))

			//region 4-set parent assets info
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

			//region 5- validate asset
			if err = objects[index].Validate(); err != nil {
				fmt.Printf(fmt.Sprintf("Error validating asset : %d : %s", index+2, err))
				return err
			}
			//endregion

			//region 6- update
			err = updateAssetEntity(fetchedAsset, objects[index], token)
			if err != nil {
				fmt.Printf(fmt.Sprintf("Error updating asset: %d : %s", index+2, err))
				return err
			}
			//endregion

			//region 7- assign asset to parent asset
			_, err = cervello.AssignAssetToAsset(objects[index].GetGlobalId(), objects[index].GetParentAssetId(), token)
			if err != nil {
				_ = cervello.DeleteAsset(objects[index].GetGlobalId(), token)
				fmt.Printf(fmt.Sprintf("error assign line no: %d to the parent asset: %s", index+2, err))
				return errors.New(fmt.Sprintf("error assign line no: %d to the parent asset: %s", index+2, err))
			}
			//endregion

			go common.PublishAuditLog("Update", entityName, objects[index].GetGlobalId(), objects[index])
		}

	}

	return nil
}

func checkUniqueAssetFeatures(obj AssetInterface) error {
	//if result, _ := common.CheckUniqueAssetField("featureId", obj.GetFeatureId(), obj.GetLayerName()); !result {
	//	return errors.New("not unique featureId")
	//}
	if result, _ := common.CheckUniqueAssetField("integrationId", obj.GetReferenceName(), ""); !result {
		return errors.New("not unique integrationId")
	}
	return nil
}

func updateAssetEntity(fetchedAsset cervello.Asset, obj AssetInterface, token string) error {
	// 1- check unique features
	//if fetchedAsset.CustomFields["featureId"] != obj.GetFeatureId() {
	//	if result, _ := common.CheckUniqueAssetField("featureId", obj.GetFeatureId(), obj.GetLayerName()); !result {
	//		return errors.New("not unique featureId")
	//	}
	//}
	if fetchedAsset.CustomFields["integrationId"] != obj.GetReferenceName() {
		if result, _ := common.CheckUniqueAssetField("integrationId", obj.GetReferenceName(), ""); !result {
			return errors.New("not unique integrationId")
		}
	}

	// 2- migrate to cervello asset
	asset, err := MigrateToCervelloAsset(obj)
	if err != nil {
		return errors.New(fmt.Sprintf("Error migrating asset : %s", err))
	}

	// 3- update asset
	err = cervello.UpdateAsset(asset.ID, asset, token)
	if err != nil {
		return err
	}

	return nil
}

func deleteAssets(csvLines [][]string, keysMap map[string]int, objects []AssetInterface, entityName string, action string, token string) error {

	//region 0- get existing entities and map them

	existingAssetsMap := map[string]cervello.Asset{}
	existingAssets, err := cervello.GetAssetsByAssetType(fmt.Sprintf("\"%s\"", objects[0].GetAssetType()), cervello.QueryParams{
		PaginationObj: cervello.Pagination{
			PageNumber: 1,
			PageSize:   99999999999,
		},
	}, token)
	if err != nil {
		return errors.New("error fetching assets")
	}
	for _, asset := range existingAssets {
		existingAssetsMap[asset.ID] = asset
	}
	//endregion

	//region 1- delete all objects
	if action == common.DeleteAction {
		for index, asset := range existingAssets {
			fmt.Printf(fmt.Sprintf("deleting line no: %d", index+2))
			err := cervello.DeleteAsset(asset.ID, token)
			if err != nil {
				return errors.New(fmt.Sprintf("error deleting line no: %d", index+2))
			}
		}
		fmt.Printf(fmt.Sprintf("all %s deleted successfully", entityName))
		return nil
	}

	//region 2- migrate from csv lines

	if err != nil {
		return err
	}
	newAssetsMap := map[string]AssetInterface{}
	for index, line := range csvLines {
		common.SleepExecution()
		fmt.Printf(fmt.Sprintf("Reading Line: %d", index+2))
		err = objects[index].MigrateFromCsvLine(line, keysMap)
		if err != nil {
			return errors.New(fmt.Sprintf("Error reading line: %d : %s", index+2, err))
		}
		newAssetsMap[objects[index].GetGlobalId()] = objects[index]
	}
	//endregion

	//region 3- delete Other assetss
	if action == common.DeleteOthersAction {
		for id := range existingAssetsMap {
			common.SleepExecution()
			if newAssetsMap[id] == nil {
				fmt.Printf(fmt.Sprintf("deleting asset with id: %s", id))
				err = cervello.DeleteAsset(id, token)
				if err != nil {
					return errors.New("error deleting asset")
				}
			}
		}
		fmt.Printf(fmt.Sprintf("other %s assets deleted successfully", entityName))
		return nil
	}
	//endregion

	//region 4- delete csv assets
	if action == common.DeleteCsvAction {
		for id := range newAssetsMap {
			common.SleepExecution()
			fmt.Printf(fmt.Sprintf("deleting asset with id: %s", id))
			err = cervello.DeleteAsset(id, token)
			if err != nil {
				return errors.New("error deleting asset")
			}

		}
		fmt.Printf(fmt.Sprintf("csv %s assets deleted successfully", entityName))
		return nil
	}
	//endregion

	return nil
}
