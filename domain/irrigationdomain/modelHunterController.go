package irrigationdomain

import (
	"databaselineservice/domain/common"
	"databaselineservice/domain/crudfunctions"
	"databaselineservice/domain/infrastructuredomain"
	"databaselineservice/sdk/cervello"
	"errors"
	"fmt"
	"strconv"
)

type HunterController struct {
	GlobalId        string                 `json:"globalId"`
	Name            string                 `json:"name"`
	IntegrationId   string                 `json:"integrationId"`
	Model           string                 `json:"model"`
	Brand           string                 `json:"brand"`
	MacAddress      string                 `json:"mac"`
	LayerName       string                 `json:"layerName"`
	LayerId         float64                `json:"layerId"`
	LayerType       string                 `json:"layerType"`
	IP              string                 `json:"ip"`
	Port            int64                  `json:"port"`
	AreaId          string                 `json:"areaId"`
	AreaNameEnglish string                 `json:"areaNameEnglish"`
	AreaNameArabic  string                 `json:"areaName"`
	AreaLayerId     float64                `json:"areaLayerId"`
	CityId          string                 `json:"cityId"`
	CityNameEnglish string                 `json:"cityNameEnglish"`
	CityNameArabic  string                 `json:"cityName"`
	CityLayerId     float64                `json:"cityLayerId"`
	X               float64                `json:"x"`
	Y               float64                `json:"y"`
	CreatedAt       string                 `json:"createdAt,omitempty"`
	UpdatedAt       string                 `json:"updatedAt,omitempty"`
	IsMaster        bool                   `json:"isMaster"`
	AdditionalInfo  map[string]interface{} `json:"additionalInfo"`
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
	return []string{deviceStateTag,
		"Hunter",
		"irr_controller",
		"irrigation",
		fmt.Sprintf("%s_alarms", thisObj.Name),
		thisObj.AreaNameEnglish,
		thisObj.AreaNameArabic}
}

func (thisObj *HunterController) SetParentAssetInfo(parentAsset cervello.Asset) error {
	area, err := infrastructuredomain.MigrateCervelloAssetToArea(parentAsset)
	if err != nil {
		return errors.New("error fetching parent floorPart: " + err.Error())
	}

	thisObj.CityId = area.CityId
	thisObj.CityNameEnglish = area.CityNameEnglish
	thisObj.CityNameArabic = area.CityNameArabic
	thisObj.AreaId = area.GlobalId
	thisObj.AreaNameEnglish = area.NameEnglish
	thisObj.AreaNameArabic = area.NameArabic
	thisObj.CityLayerId = area.CityLayerId
	thisObj.AreaLayerId = area.LayerId

	return nil
}

func (thisObj *HunterController) GetLayerName() string {
	return thisObj.LayerName
}

func (thisObj *HunterController) GetLayerId() float64 {
	return thisObj.LayerId
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
		Configuration: []cervello.ModbusConfiguration{
			{
				Address: 7001,
				Mapping: map[string]map[string]string{
					"7001": {
						"key":  "register7001",
						"type": "TELEMETRY",
					},
				},
				SlaveID:       1,
				Quantity:      1,
				Sequence:      1,
				OperationCode: 3,
			},
		},
		Schedule: cervello.ModbusConfigurationSchedule{
			Interval: 10,
			TimeUnit: "Second",
			Timezone: "Africa/Cairo",
		},
	}
}

func (thisObj *HunterController) MigrateFromCsvLine(csvLine []string, keysMap map[string]int) error {
	var err error
	thisObj.GlobalId = csvLine[keysMap["globalid"]]
	thisObj.IntegrationId = csvLine[keysMap["integrationuuid"]]
	thisObj.Name = csvLine[keysMap["name"]]
	thisObj.CityId = csvLine[keysMap["city_uuid"]]
	thisObj.AreaId = csvLine[keysMap["area_uuid"]]
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
	thisObj.LayerId, err = strconv.ParseFloat(csvLine[keysMap["layerid"]], 64)
	if err != nil {
		return err
	}
	thisObj.LayerName = csvLine[keysMap["layername"]]

	thisObj.Brand = csvLine[keysMap["brand"]]
	thisObj.Model = csvLine[keysMap["modelno"]]
	thisObj.IP = csvLine[keysMap["ip"]]
	thisObj.Port, err = strconv.ParseInt(csvLine[keysMap["port"]], 10, 64)
	if err != nil {
		return err
	}
	thisObj.MacAddress = csvLine[keysMap["mac"]]
	thisObj.LayerType = "point"
	thisObj.AdditionalInfo = common.SetupAdditionalInfo(keysMap, thisObj.GetEssentialKeys(), csvLine)
	thisObj.Name = thisObj.Name + "_" + thisObj.MacAddress
	thisObj.IsMaster, err = strconv.ParseBool(csvLine[keysMap["is_master"]])
	if err != nil {
		return err
	}

	return err
}

func (thisObj *HunterController) GetEssentialKeys() []string {
	return []string{
		"id",
		"globalid",
		"integrationuuid",
		"city_uuid",
		"area_uuid",
		"name",
		"ip",
		"layerid",
		"layername",
		"x",
		"y",
		"port",
		"brand",
		"mac",
		"is_master",
	}
}

func (thisObj *HunterController) GetNonDuplicatingKeys() []string {
	return []string{
		"id",
		"globalid",
		"integrationuuid",
		"name",
		"mac",
		"ip",
	}
}

func (thisObj *HunterController) GetParentAssetKey() string {
	return "area_uuid"
}

func (thisObj *HunterController) GetParentGatewayKey() string {
	return common.EmptyField
}
