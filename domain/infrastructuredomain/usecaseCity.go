package infrastructuredomain

import (
	"databaselineservice/domain/common"
	"databaselineservice/sdk/cervello"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type City struct {
	GlobalId    string  `json:"globalId"`
	NameEnglish string  `json:"name"`
	NameArabic  string  `json:"nameArabic"`
	Area        float64 `json:"area"`
	LayerType   string  `json:"layerType"`
	CreatedAt   string  `json:"createdAt,omitempty"`
	UpdatedAt   string  `json:"updatedAt,omitempty"`
}

func MigrateCityToCervelloAsset(city City) (cervello.Asset, error) {
	toMap, err := common.StructToMap(city)
	if err != nil {
		return cervello.Asset{}, err
	}

	return cervello.Asset{
		ID:            city.GlobalId,
		AssetType:     "City",
		Name:          city.NameEnglish,
		ReferenceName: "CITY",
		CustomFields:  toMap,
	}, nil
}

//
//func MigrateCityToCervelloDevice(city City) (cervello.Device, error) {
//	toMap, err := common.StructToMap(city)
//	if err != nil {
//		return cervello.Device{}, err
//	}
//
//	return cervello.Device{
//		ID:           city.GlobalId,
//		Tags:         []string{"citylight"},
//		Name:         city.Name + " Light",
//		DeviceType:   cervello.DeviceTypeStandalone,
//		CustomFields: toMap,
//	}, nil
//}

func MigrateCervelloAssetToCity(asset cervello.Asset) (City, error) {
	result := City{}

	jsonResult, err := json.Marshal(asset.CustomFields)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(jsonResult, &result)
	if err != nil {
		return result, err
	}

	result.GlobalId = asset.ID
	result.NameEnglish = asset.Name
	result.CreatedAt = asset.CreatedAt
	result.UpdatedAt = asset.UpdatedAt

	return result, err
}

func (city *City) Validate() error {

	if city.GlobalId == "" {
		return errors.New("ID is missing")
	}

	if !common.IsValidUUID(city.GlobalId) {
		return errors.New("ID must be uuid")
	}

	if city.LayerType == "" {
		return errors.New("LayerType is missing")
	}

	return nil
}

func (city *City) GetAssetType() string {
	return "cityRolesTest"
}

func MigrateCityFromCSVLine(csvLine []string, keysMap map[string]int) (City, error) {
	var err error = nil
	result := City{}

	result.GlobalId = csvLine[keysMap["globalId"]]
	result.NameEnglish = csvLine[keysMap["nameEn"]]
	result.NameArabic = csvLine[keysMap["nameAr"]]
	result.NameArabic = strings.Replace(result.NameArabic, " ", "_", -1)
	result.NameArabic = strings.Replace(result.NameArabic, ".", "", -1)
	result.LayerType = "Polygon"
	result.Area, err = strconv.ParseFloat(csvLine[keysMap["area"]], 64)
	if err != nil {
		result.Area = 0
	}

	return result, nil
}

func (thisObj *City) GetEssentialKeys() []string {
	return []string{
		"globalid",
		"nameEn",
		"nameAr",
	}
}

func UseCaseCities(csvLines [][]string, keysMap map[string]int, action string) (string, error) {
	for index, line := range csvLines {
		fmt.Println("Reading Line: %d", index+1)
		city, err := MigrateCityFromCSVLine(line, keysMap)
		if err != nil {
			return "", err
		}
		//fmt.Println(city)
		if err := city.Validate(); err != nil {
			return fmt.Sprintf("error reading Line: %d", index+1), err
		}
		if action == "validate" {
			fmt.Println("city no: %d is valid", index+1)
			continue
		}
		cityAsset, err := MigrateCityToCervelloAsset(city)
		if err != nil {
			return "", err
		}

		//cityDevice, err := MigrateCityToCervelloDevice(city)
		//if err != nil {
		//	return "", err
		//}

		if fetchedCity, err := cervello.GetAssetByID(city.GlobalId, ""); err != nil || fetchedCity == nil || fetchedCity.ID == "" {
			fmt.Println("creating city no: %d", index+1)
			_, err = cervello.CreateAsset(cityAsset, "")
			if err != nil {
				return fmt.Sprintf("error creating city no: %d", index+1), err
			}
			//_, err = cervello.CreateDevice(cityDevice, "")
			//if err != nil {
			//	_ = cervello.DeleteAsset(city.GlobalId, "")
			//	return fmt.Sprintf("error creating city no: %d", index+1), err
			//}

			go common.PublishAuditLog("Create", "City", city.GlobalId, city)
		} else {
			fmt.Println("updating city no: %d", index+1)
			_, err = UpdateCity(cityAsset)
			if err != nil {
				return fmt.Sprintf("error updating city no: %d", index+1), err
			}

			go common.PublishAuditLog("Update", "City", city.GlobalId, city)
		}
	}
	if action == "validate" {
		return "cities are validated successfully", nil
	}
	return "cities are created successfully", nil
}
