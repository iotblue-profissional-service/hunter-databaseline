package infrastructuredomain

import (
	"databaselineservice/domain/common"
	"databaselineservice/domain/crudfunctions"
	"databaselineservice/sdk/cervello"
	"strconv"
	"strings"
)

type Geometry struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

type Area struct {
	GlobalId         string   `json:"globalId"`
	NameEnglish      string   `json:"name"`
	NameArabic       string   `json:"name_Arabic"`
	PhaseId          string   `json:"phaseId"`
	PhaseNameEnglish string   `json:"phaseNameEnglish"`
	PhaseNameArabic  string   `json:"phaseName"`
	CityId           string   `json:"cityId"`
	CityNameEnglish  string   `json:"cityName"`
	CityNameArabic   string   `json:"cityNameArabic"`
	CreatedAt        string   `json:"createdAt,omitempty"`
	UpdatedAt        string   `json:"updatedAt,omitempty"`
	Area             float64  `json:"area"`
	LayerType        string   `json:"layerType"`
	Coordinates      Geometry `json:"Geometry`
}

func (thisObj *Area) GetGlobalId() string {
	return thisObj.GlobalId
}

func (thisObj *Area) GetName() string {
	return thisObj.NameEnglish
}

func (thisObj *Area) GetModel() string {
	return common.EmptyField
}

func (thisObj *Area) ValidateModel() error {
	return nil
}

func (thisObj *Area) GetReferenceName() string {
	return thisObj.NameEnglish
}

func (thisObj *Area) GetClientId() string {
	return ""
}

func (thisObj *Area) GetIP() string {
	return common.EmptyField
}

func (thisObj *Area) GetFeatureId() string {
	return common.EmptyField
}

func (thisObj *Area) GetParentAssetId() string {
	return thisObj.CityId
}

func (thisObj *Area) GetParentGatewayId() string {
	return common.EmptyField
}

func (thisObj *Area) GetDeviceType() string {
	return cervello.DeviceTypeStandalone
}

func (thisObj *Area) GetTags() []string {
	return []string{
		"Hunter",
		"asset",
		"area",
		"irrigation",
	}
}

func (thisObj *Area) GetSearchTag() string {
	return "area"
}

func (thisObj *Area) GetAssetType() string {
	return "area"
}

func (thisObj *Area) SetParentGatewayInfo(_ cervello.Device) error {
	return nil
}

func (thisObj *Area) SetParentAssetInfo(parentAsset cervello.Asset) error {
	city, err := MigrateCervelloAssetToCity(parentAsset)
	if err != nil {
		return err
	}
	thisObj.CityId = city.GlobalId
	thisObj.CityNameEnglish = city.NameEnglish
	thisObj.CityNameArabic = city.NameArabic
	return nil
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
	thisObj.PhaseId = ""
	thisObj.NameEnglish = csvLine[keysMap["nameEn"]]
	thisObj.NameArabic = csvLine[keysMap["nameAr"]]
	thisObj.NameArabic = strings.Replace(thisObj.NameArabic, " ", "_", -1)
	thisObj.NameArabic = strings.Replace(thisObj.NameArabic, ".", "", -1)
	thisObj.LayerType = "Polygon"
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
		"nameEn",
		"nameAr",
	}
}

func (thisObj *Area) GetParentAssetKey() string {
	return "cityId"
}

func (thisObj *Area) GetParentGatewayKey() string {
	return common.EmptyField
}

func (thisObj *Area) GetMac() string {
	return common.EmptyField
}

func (thisObj *Area) GetParentGatewayKeyGetMac() string {
	return common.EmptyField
}

func (thisObj *Area) GetCommunicationProtocolConfiguration() (string, interface{}) {
	return common.EmptyField, nil
}
