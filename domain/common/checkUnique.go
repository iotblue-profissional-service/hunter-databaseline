package common

import (
	"databaselineservice/sdk/cervello"
	"fmt"
)

func CheckUniqueDeviceField(fieldName string, value string, layerName string) (bool, int) {
	filterQuery := map[string]interface{}{}
	NumberOfFilters := 0
	filterQuery[fmt.Sprintf("filters[%d][key]", NumberOfFilters)] = "customFields." + fieldName
	filterQuery[fmt.Sprintf("filters[%d][operator]", NumberOfFilters)] = "eq"
	filterQuery[fmt.Sprintf("filters[%d][value]", NumberOfFilters)] = value
	NumberOfFilters += 1

	if layerName != "" {
		filterQuery[fmt.Sprintf("filters[%d][key]", NumberOfFilters)] = "customFields.layerName"
		filterQuery[fmt.Sprintf("filters[%d][operator]", NumberOfFilters)] = "eq"
		filterQuery[fmt.Sprintf("filters[%d][value]", NumberOfFilters)] = layerName
		NumberOfFilters += 1
	}
	cervelloQuery := cervello.QueryParams{
		Custom: filterQuery,
		PaginationObj: cervello.Pagination{
			PageNumber: 1,
			PageSize:   1,
		},
	}
	filteredList := make([]cervello.Device, 0)
	filteredList, _ = cervello.GetOrgDevicesFiltered(cervelloQuery, "")
	if len(filteredList) > 0 {
		return false, len(filteredList)
	}
	return true, len(filteredList)
}

func CheckUniqueAssetField(fieldName string, value string, tag string) (bool, int) {
	filterQuery := map[string]interface{}{}
	NumberOfFilters := 0
	filterQuery[fmt.Sprintf("filters[%d][key]", NumberOfFilters)] = "customFields." + fieldName
	filterQuery[fmt.Sprintf("filters[%d][operator]", NumberOfFilters)] = "eq"
	filterQuery[fmt.Sprintf("filters[%d][value]", NumberOfFilters)] = value
	NumberOfFilters += 1

	if tag != "" {
		filterQuery[fmt.Sprintf("filters[%d][key]", NumberOfFilters)] = "tags"
		filterQuery[fmt.Sprintf("filters[%d][operator]", NumberOfFilters)] = "contains"
		filterQuery[fmt.Sprintf("filters[%d][value]", NumberOfFilters)] = tag
		NumberOfFilters += 1
	}

	cervelloQuery := cervello.QueryParams{
		Custom: filterQuery,
		PaginationObj: cervello.Pagination{
			PageNumber: 1,
			PageSize:   1,
		},
	}
	filteredList := make([]cervello.Asset, 0)
	filteredList, _ = cervello.GetAssetsFlitered(cervelloQuery, "")
	if len(filteredList) > 0 {
		return false, len(filteredList)
	}
	return true, len(filteredList)
}
