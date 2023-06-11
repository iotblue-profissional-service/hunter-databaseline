package irrigationdomain

import (
	"databaselineservice/domain/common"
	"databaselineservice/domain/crudfunctions"
	"databaselineservice/domain/infrastructuredomain"
	"databaselineservice/domain/irrigationdomain/modbusConfig"
	"databaselineservice/sdk/cervello"
	"errors"
	"strconv"
	"strings"
)

type WeatherStation struct {
	GlobalId       string                 `json:"globalId"`
	Type           string                 `json:"type"`
	Name           string                 `json:"name"`
	IntegrationId  string                 `json:"integrationId"`
	MacAddress     string                 `json:"mac"`
	LayerType      string                 `json:"layerType"`
	IP             string                 `json:"ip"`
	Port           int64                  `json:"port"`
	AreaId         string                 `json:"areaId"`
	AreaName       string                 `json:"areaName"`
	AreaNameArabic string                 `json:"areaNameArabic"`
	CityId         string                 `json:"cityId"`
	CityName       string                 `json:"cityName"`
	CityNameArabic string                 `json:"cityNameArabic"`
	X              float64                `json:"x"`
	Y              float64                `json:"y"`
	CreatedAt      string                 `json:"createdAt,omitempty"`
	UpdatedAt      string                 `json:"updatedAt,omitempty"`
	AdditionalInfo map[string]interface{} `json:"additionalInfo"`
}

func (thisObj *WeatherStation) GetMac() string {
	return thisObj.MacAddress
	//return common.EmptyField
}

func (thisObj *WeatherStation) GetHost() string {
	return thisObj.IP
}

func (thisObj *WeatherStation) GetPort() int64 {
	return thisObj.Port
}

func (thisObj *WeatherStation) GetGlobalId() string {
	return thisObj.GlobalId
}

func (thisObj *WeatherStation) GetName() string {
	return thisObj.Name
}

func (thisObj *WeatherStation) ValidateModel() error {
	// un implemented until i hava a list of models
	return nil
}

func (thisObj *WeatherStation) GetReferenceName() string {
	return thisObj.IntegrationId
}

func (thisObj *WeatherStation) GetClientId() string {
	return thisObj.IntegrationId
}

func (thisObj *WeatherStation) GetIP() string {
	return thisObj.IP
}

func (thisObj *WeatherStation) GetParentAssetId() string {
	return thisObj.AreaId
}

func (thisObj *WeatherStation) GetParentGatewayId() string {
	return common.EmptyField
}

func (thisObj *WeatherStation) GetDeviceType() string {
	return cervello.DeviceTypeStandalone
}

func (thisObj *WeatherStation) GetTags() []string {
	deviceStateTag := common.MockDevice
	if common.IsPhysicalDevice {
		deviceStateTag = common.GisDevice
	}
	areaTag := strings.Replace(thisObj.AreaName, " ", "_", -1)
	areaTag = strings.Replace(areaTag, ".", "", -1)
	return []string{deviceStateTag,
		"Hunter",
		"weather_station",
		"irrigation",
		areaTag,
	}
}

func (thisObj *WeatherStation) SetParentAssetInfo(parentAsset cervello.Asset) error {
	area, err := infrastructuredomain.MigrateCervelloAssetToArea(parentAsset)
	if err != nil {
		return errors.New("error fetching parent area: " + err.Error())
	}

	thisObj.CityId = area.CityId
	thisObj.CityName = area.CityNameEnglish
	thisObj.CityNameArabic = area.CityNameArabic
	thisObj.AreaId = area.GlobalId
	thisObj.AreaName = area.NameEnglish
	thisObj.AreaNameArabic = area.NameArabic

	return nil
}

func (thisObj *WeatherStation) GetLayerType() string {
	return thisObj.LayerType
}

func (thisObj *WeatherStation) SetParentGatewayInfo(_ cervello.Device) error {
	return nil
}

func (thisObj *WeatherStation) Validate() error {
	if thisObj.MacAddress == "" {
		return errors.New("mac address is missing")
	}

	if !common.IsValidMacAddress(thisObj.MacAddress) {
		return errors.New("invalid mac address")
	}

	return crudfunctions.ValidateModbusDeviceEntity(thisObj)
}

func (thisObj *WeatherStation) GetSearchTag() string {
	return "weather_station"
}

func (thisObj *WeatherStation) GetModbusConfig() *cervello.ModbusDeviceConfig {
	return &cervello.ModbusDeviceConfig{
		Configuration: modbusConfig.WeatherStationConfig[:],
		Schedule: cervello.ModbusConfigurationSchedule{
			Interval: 10,
			TimeUnit: "Second",
			Timezone: "Africa/Cairo",
		},
	}
}

func (thisObj *WeatherStation) MigrateFromCsvLine(csvLine []string, keysMap map[string]int) error {
	var err error
	thisObj.GlobalId = csvLine[keysMap["globalId"]]
	thisObj.IntegrationId = csvLine[keysMap["integrationId"]]
	thisObj.Name = csvLine[keysMap["name"]]
	thisObj.AreaId = csvLine[keysMap["areaId"]]
	if err != nil {
		return err
	}
	thisObj.X, err = strconv.ParseFloat(csvLine[keysMap["x"]], 64)
	if err != nil {
		return err
	}
	thisObj.Y, err = strconv.ParseFloat(csvLine[keysMap["y"]], 64)
	if err != nil {
		return err
	}
	thisObj.IP = csvLine[keysMap["ip"]]
	thisObj.Port, err = strconv.ParseInt(csvLine[keysMap["port"]], 10, 64)
	if err != nil {
		return err
	}
	thisObj.MacAddress = csvLine[keysMap["mac"]]
	thisObj.LayerType = "point"
	thisObj.Type = "weatherStation"
	thisObj.AdditionalInfo = common.SetupAdditionalInfo(keysMap, thisObj.GetEssentialKeys(), csvLine)

	return err
}

func (thisObj *WeatherStation) GetEssentialKeys() []string {
	return []string{
		"globalId",
		"integrationId",
		"areaId",
		"name",
		"ip",
		"x",
		"y",
		"port",
		"mac",
	}
}

func (thisObj *WeatherStation) GetNonDuplicatingKeys() []string {
	return []string{
		"globalId",
		"integrationId",
		"name",
		"mac",
		"ip",
	}
}

func (thisObj *WeatherStation) GetParentAssetKey() string {
	return "areaId"
}

func (thisObj *WeatherStation) GetParentGatewayKey() string {
	return common.EmptyField
}
