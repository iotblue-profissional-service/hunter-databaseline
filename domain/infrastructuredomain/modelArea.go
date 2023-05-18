package infrastructuredomain

import (
	"databaselineservice/domain/common"
	"databaselineservice/domain/crudfunctions"
	"databaselineservice/sdk/cervello"
	"strconv"
	"strings"
)

type Area struct {
	GlobalId         string  `json:"globalId"`
	NameEnglish      string  `json:"name_English"`
	NameArabic       string  `json:"name"`
	PhaseId          string  `json:"phaseId"`
	PhaseNameEnglish string  `json:"phaseNameEnglish"`
	PhaseNameArabic  string  `json:"phaseName"`
	CityId           string  `json:"cityId"`
	CityNameEnglish  string  `json:"cityNameEnglish"`
	CityNameArabic   string  `json:"cityName"`
	LayerId          float64 `json:"layerId"`
	PhaseLayerID     float64 `json:"phaseLayerId"`
	CityLayerId      float64 `json:"cityLayerId"`
	CreatedAt        string  `json:"createdAt,omitempty"`
	UpdatedAt        string  `json:"updatedAt,omitempty"`
	Area             float64 `json:"area"`
	LayerType        string  `json:"layerType"`
	LayerName        string  `json:"layerName"`
}

func (thisObj *Area) GetGlobalId() string {
	return thisObj.GlobalId
}

func (thisObj *Area) GetName() string {
	return thisObj.NameArabic
}

func (thisObj *Area) GetModel() string {
	return common.EmptyField
}

func (thisObj *Area) ValidateModel() error {
	return nil
}

func (thisObj *Area) GetReferenceName() string {
	return common.EmptyField
}

func (thisObj *Area) GetFeatureId() string {
	return common.EmptyField
}

func (thisObj *Area) GetParentAssetId() string {
	return thisObj.CityId
}

func (thisObj *Area) GetAssetType() string {
	return "olympicArea"
}

func (thisObj *Area) SetParentAssetInfo(parentAsset cervello.Asset) error {
	city, err := MigrateCervelloAssetToCity(parentAsset)
	if err != nil {
		return err
	}
	thisObj.CityId = city.GlobalId
	thisObj.CityNameEnglish = city.NameEnglish
	thisObj.CityNameArabic = city.NameArabic
	thisObj.CityLayerId = city.LayerId
	return nil
}

func (thisObj *Area) GetLayerName() string {
	return thisObj.LayerName
}

func (thisObj *Area) GetLayerId() float64 {
	return thisObj.LayerId
}

func (thisObj *Area) GetLayerType() string {
	return thisObj.LayerType
}

func (thisObj *Area) Validate() error {
	return crudfunctions.ValidateAssetEntity(thisObj)
}

func (thisObj *Area) MigrateFromCsvLine(csvLine []string, keysMap map[string]int) error {
	var err error = nil
	thisObj.GlobalId = csvLine[keysMap["globalId"]]
	thisObj.CityId = csvLine[keysMap["cityId"]]
	thisObj.PhaseId = common.EmptyField
	thisObj.NameEnglish = csvLine[keysMap["nameEn"]]
	thisObj.NameArabic = csvLine[keysMap["nameAr"]]
	thisObj.NameArabic = strings.Replace(thisObj.NameArabic, " ", "_", -1)
	thisObj.NameArabic = strings.Replace(thisObj.NameArabic, ".", "", -1)
	thisObj.LayerId = common.AreaLayerId
	thisObj.LayerType = "Polygon"
	thisObj.LayerName = common.AreaLayerName
	thisObj.Area, err = strconv.ParseFloat(csvLine[keysMap["area"]], 64)
	if err != nil {
		thisObj.Area = 0
	}

	return nil
}

func (thisObj *Area) GetEssentialKeys() []string {
	return []string{
		"globalId",
		"cityId",
		"nameEn",
		"nameAr",
		"area",
	}
}

func (thisObj *Area) GetNonDuplicatingKeys() []string {
	return []string{
		"globalId",
	}
}

func (thisObj *Area) GetParentAssetKey() string {
	return "cityId"
}
