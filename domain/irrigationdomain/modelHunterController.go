package irrigationdomain

import (
	"databaselineservice/domain/common"
	"databaselineservice/domain/crudfunctions"
	"databaselineservice/domain/infrastructuredomain"
	"databaselineservice/domain/irrigationdomain/modbusConfig"
	"databaselineservice/sdk/cervello"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type HunterController struct {
	GlobalId         string                 `json:"globalId"`
	ControllerId     string                 `json:"controllerId"`
	Type             string                 `json:"type"`
	Name             string                 `json:"name"`
	IntegrationId    string                 `json:"integrationId"`
	Model            string                 `json:"model"`
	Brand            string                 `json:"brand"`
	MacAddress       string                 `json:"mac"`
	LayerType        string                 `json:"layerType"`
	IP               string                 `json:"ip"`
	Port             int64                  `json:"port"`
	AreaId           string                 `json:"areaId"`
	AreaName         string                 `json:"areaName"`
	AreaNameArabic   string                 `json:"areaNameArabic"`
	CityId           string                 `json:"cityId"`
	CityName         string                 `json:"cityName"`
	CityNameArabic   string                 `json:"cityNameArabic"`
	X                float64                `json:"x"`
	Y                float64                `json:"y"`
	CreatedAt        string                 `json:"createdAt,omitempty"`
	UpdatedAt        string                 `json:"updatedAt,omitempty"`
	StationCount     int64                  `json:"stationCount"`
	FlowSensorCount  int64                  `json:"flowSensorCount"`
	MasterValveCount int64                  `json:"masterValveCount"`
	AdditionalInfo   map[string]interface{} `json:"additionalInfo"`
}

func (thisObj *HunterController) GetMac() string {
	return thisObj.MacAddress
	//return common.EmptyField
}

func (thisObj *HunterController) GetHost() string {
	return thisObj.IP
}

func (thisObj *HunterController) GetPort() int64 {
	return thisObj.Port
}

func (thisObj *HunterController) GetGlobalId() string {
	return thisObj.GlobalId
}

func (thisObj *HunterController) GetName() string {
	return thisObj.Name
}

func (thisObj *HunterController) GetModel() string {
	return thisObj.Model
}

func (thisObj *HunterController) ValidateModel() error {
	// un implemented until i hava a list of models
	return nil
}

func (thisObj *HunterController) GetReferenceName() string {
	return thisObj.IntegrationId
}

func (thisObj *HunterController) GetClientId() string {
	return thisObj.IntegrationId
}

func (thisObj *HunterController) GetIP() string {
	return thisObj.IP
}

func (thisObj *HunterController) GetParentAssetId() string {
	return thisObj.AreaId
}

func (thisObj *HunterController) GetParentGatewayId() string {
	return common.EmptyField
}

func (thisObj *HunterController) GetDeviceType() string {
	return cervello.DeviceTypeGateWay
}

func (thisObj *HunterController) GetTags() []string {
	deviceStateTag := common.MockDevice
	if common.IsPhysicalDevice {
		deviceStateTag = common.GisDevice
	}
	areaTag := strings.Replace(thisObj.AreaName, " ", "_", -1)
	areaTag = strings.Replace(areaTag, ".", "", -1)
	return []string{deviceStateTag,
		"Hunter",
		"irr_controller",
		"irrigation",
		fmt.Sprintf("%s_alarms", thisObj.Name),
		areaTag,
	}
}

func (thisObj *HunterController) SetParentAssetInfo(parentAsset cervello.Asset) error {
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

func (thisObj *HunterController) GetLayerType() string {
	return thisObj.LayerType
}

func (thisObj *HunterController) SetParentGatewayInfo(_ cervello.Device) error {
	return nil
}

func (thisObj *HunterController) Validate() error {
	if thisObj.MacAddress == "" {
		return errors.New("mac address is missing")
	}

	if !common.IsValidMacAddress(thisObj.MacAddress) {
		return errors.New("invalid mac address")
	}

	return crudfunctions.ValidateModbusDeviceEntity(thisObj)
}

func (thisObj *HunterController) GetSearchTag() string {
	return "irr_controller"
}

func (thisObj *HunterController) GetModbusConfig() *cervello.ModbusDeviceConfig {
	return &cervello.ModbusDeviceConfig{
		Configuration: modbusConfig.HunterControllerConfig[:],
		Schedule: cervello.ModbusConfigurationSchedule{
			Interval: 5,
			TimeUnit: "Second",
			Timezone: "Africa/Cairo",
		},
	}
}

func (thisObj *HunterController) MigrateFromCsvLine(csvLine []string, keysMap map[string]int) error {
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
	thisObj.Brand = csvLine[keysMap["brand"]]
	thisObj.Model = csvLine[keysMap["model"]]
	thisObj.IP = csvLine[keysMap["ip"]]
	thisObj.Port, err = strconv.ParseInt(csvLine[keysMap["port"]], 10, 64)
	if err != nil {
		return err
	}
	thisObj.ControllerId = thisObj.GlobalId
	thisObj.MacAddress = csvLine[keysMap["mac"]]
	thisObj.LayerType = "point"
	thisObj.Type = "controller"
	thisObj.StationCount, err = strconv.ParseInt(csvLine[keysMap["stationCount"]], 10, 64)
	if err != nil {
		return err
	}
	thisObj.FlowSensorCount, err = strconv.ParseInt(csvLine[keysMap["flowSensorCount"]], 10, 64)
	if err != nil {
		return err
	}
	thisObj.MasterValveCount, err = strconv.ParseInt(csvLine[keysMap["masterValveCount"]], 10, 64)
	if err != nil {
		return err
	}
	thisObj.AdditionalInfo = common.SetupAdditionalInfo(keysMap, thisObj.GetEssentialKeys(), csvLine)

	return err
}

func (thisObj *HunterController) GetEssentialKeys() []string {
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
		"stationCount",
		"flowSensorCount",
		"masterValveCount",
		"brand",
		"model",
	}
}

func (thisObj *HunterController) GetNonDuplicatingKeys() []string {
	return []string{
		"globalId",
		"integrationId",
		"name",
		"mac",
		"ip",
	}
}

func (thisObj *HunterController) GetParentAssetKey() string {
	return "areaId"
}

func (thisObj *HunterController) GetParentGatewayKey() string {
	return common.EmptyField
}
