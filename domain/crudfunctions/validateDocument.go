package crudfunctions

import (
	"databaselineservice/domain/common"
	"databaselineservice/sdk/cervello"
	"errors"
	"fmt"
)

type ParentsData struct {
	ParentAssetType        string
	ParentAssetKey         string
	ParentGatewaySearchTag string
	ParentGatewayKey       string
}

type ValidatingObject interface {
	GetNonDuplicatingKeys() []string
	MigrateFromCsvLine(csvLine []string, keysMap map[string]int) error
	Validate() error
}

func ValidateDocument(csvLines [][]string,
	keysMap map[string]int,
	ColumnNames []string, validatingObjects []ValidatingObject, parentsData ParentsData, token string) error {

	errCount := 0
	// 1- validate all keys exist
	err := common.ValidateAllColumnsExist(keysMap, ColumnNames)
	if err != nil {
		fmt.Printf(fmt.Sprintf("%v", err))
		errCount += 1
	}
	if errCount > 0 {
		// if there are missing columns don`t continue
		return errors.New("not valid document")
	}

	// 2- validate non duplicating keys
	for _, key := range validatingObjects[0].GetNonDuplicatingKeys() {
		type countMap map[string]int
		count := map[string]countMap{}
		for index, line := range csvLines {
			variable := line[keysMap[key]]
			if count[variable] == nil {
				count[variable] = make(countMap)
			}
			count[variable]["counts"] += 1
			if count[variable]["counts"] > 1 {
				fmt.Printf(fmt.Sprintf("there are duplications in column %s with row %d and row %d", key, index+2, count[variable]["firstLine"]))
				errCount += 1
			}
			if count[variable]["firstLine"] == 0 {
				count[variable]["firstLine"] = index + 2
			}
		}

	}

	// 3- validate data
	for index, line := range csvLines {
		err = validatingObjects[index].MigrateFromCsvLine(line, keysMap)
		if err != nil {
			fmt.Printf(fmt.Sprintf("there is error in line %d :%v", index+2, err))
			errCount += 1
		}

		err = validatingObjects[index].Validate()
		if err != nil {
			fmt.Printf(fmt.Sprintf("there is error in line %d :%v", index+2, err))
			errCount += 1
		}
	}

	// 4- validate parents
	paginationObj := cervello.Pagination{PageSize: 999999, PageNumber: 0}
	if parentsData.ParentAssetType != common.EmptyField {
		parentAssets, err := cervello.GetAssetsByAssetType(fmt.Sprintf("\"%s\"", parentsData.ParentAssetType), cervello.QueryParams{PaginationObj: paginationObj}, token)
		if err != nil {
			fmt.Printf(fmt.Sprintf("error fetching parent assets from the database: %v", err))
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

		for index, line := range csvLines {
			assetKey := parentsData.ParentAssetKey
			parentAssetId := line[keysMap[assetKey]]
			if parentAssetMap[parentAssetId].ID == "" {
				fmt.Printf(fmt.Sprintf("%s doesnt exist in database for line %d", assetKey, index+2))
				errCount += 1
			}
		}
	}
	if parentsData.ParentGatewaySearchTag != common.EmptyField {
		parentGatewayMap := make(map[string]cervello.Device)
		parentGateways, err := cervello.GetOrgDevicesFiltered(cervello.QueryParams{
			PaginationObj: paginationObj,
			Filters: []cervello.Filter{{
				Key:   "tags",
				Op:    "contains",
				Value: fmt.Sprintf("\"%s\"", parentsData.ParentGatewaySearchTag),
			}},
		}, token)
		if (err != nil) || (len(parentGateways) == 0) {
			fmt.Printf("no parent gateways in the database")
			return errors.New("no parent gateways in the database")
		}
		for _, parent := range parentGateways {
			parentGatewayMap[parent.ID] = parent
		}
		for index, line := range csvLines {
			gatewayKey := parentsData.ParentGatewayKey
			parentGatewayId := line[keysMap[gatewayKey]]
			if parentGatewayMap[parentGatewayId].ID == "" {
				fmt.Printf(fmt.Sprintf("%s doesnt exist in database for line %d", gatewayKey, index+2))
				errCount += 1
			}
		}
	}

	if errCount > 0 {
		return errors.New("not valid document")
	}
	return nil
}
