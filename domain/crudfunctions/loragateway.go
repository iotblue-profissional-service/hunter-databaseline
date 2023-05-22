package crudfunctions

import (
	"databaselineservice/domain/common"
	"databaselineservice/sdk/cervello"
	"errors"
	"fmt"
)

func UseCaseLoraGatewayEntity(csvLines [][]string, keysMap map[string]int, objects []LoraGatewayInterface, parentAssetType string, entityName string, action string, token string) error {
	//region 1- migrate from csv lines
	ColumnNames := objects[0].GetEssentialKeys()
	err := common.ValidateAllColumnsExist(keysMap, ColumnNames)
	if err != nil {
		return err
	}
	for index, line := range csvLines {
		fmt.Println(fmt.Sprintf("Reading Line: %d", index+2))
		err = objects[index].MigrateFromCsvLine(line, keysMap)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error reading line: %d : %s", index+2, err))
			return err
		}
	}
	//endregion

	//region 2- fetch parent assets
	paginationObj := cervello.Pagination{PageSize: 999999, PageNumber: 0}
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

	for index := range objects {
		common.SleepExecution()

		// create device (or update if it exists)
		if fetchedDevice, _ := cervello.LoraService.GetGatewayByID(objects[index].GetGlobalId(), token); fetchedDevice == nil && action != common.UpdateOnlyAction {
			fmt.Println(fmt.Sprintf("creating line no: %d", index+2))

			//region 3- set parent assets info
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

			//region 4- validate device
			if err := objects[index].Validate(); err != nil {
				fmt.Println(fmt.Sprintf("Error validating device : %d : %s", index+2, err))
				return err
			}
			//endregion

			//region 5- create
			loraGateway, err := MigrateToCervelloLoraGateway(objects[index])
			if err != nil {
				return err
			}

			//  create lora gateway
			_, err = cervello.LoraService.CreateGateway(loraGateway, token)
			if err != nil {
				return err
			}
			//endregion

			go common.PublishAuditLog("Create", entityName, objects[index].GetGlobalId(), objects[index])
		} else {
			if action == common.CreateOnlyAction || fetchedDevice.ID == "" {
				continue
			}
			fmt.Println(fmt.Sprintf("updating line no: %d", index+2))

			//region 3- set parent assets info
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

			//region 4- validate device
			if err := objects[index].Validate(); err != nil {
				fmt.Println(fmt.Sprintf("Error validating device : %d : %s", index+2, err))
				return err
			}
			//endregion

			//region 5- update device
			err := updateLoraGatewayEntity(*fetchedDevice, objects[index], token)
			if err != nil {
				fmt.Println(fmt.Sprintf("Error updating device: %d : %s", index+2, err))
				return err
			}
			//endregion

			go common.PublishAuditLog("Update", entityName, objects[index].GetGlobalId(), objects[index])
		}

	}

	return nil
}

func updateLoraGatewayEntity(fetchedDevice cervello.LoraGateway, obj LoraGatewayInterface, token string) error {

	// 1- to lora gateway
	loraGateway, err := MigrateToCervelloLoraGateway(obj)
	if err != nil {
		return err
	}

	// 2- update lora gateway
	_, err = cervello.LoraService.UpdateGateway(fetchedDevice.ID, loraGateway, token)
	if err != nil {
		return errors.New("cant update device: " + err.Error())
	}

	return nil
}
