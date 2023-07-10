package irrigationdomain

import (
	"databaselineservice/domain/common"
	"databaselineservice/domain/crudfunctions"
	"databaselineservice/domain/infrastructuredomain"
	"databaselineservice/sdk/cervello"
	"errors"
	"strconv"
)

type HunterFlowSensor struct {
	GlobalId        string                 `json:"globalId"`
	Name            string                 `json:"name"`
	Type            string                 `json:"type"`
	LayerType       string                 `json:"layerType"`
	ControllerName  string                 `json:"controllerName"`
	ControllerId    string                 `json:"controllerId"`
	AreaId          string                 `json:"areaId"`
	AreaNameEnglish string                 `json:"areaNameEnglish"`
	AreaNameArabic  string                 `json:"areaName"`
	CityId          string                 `json:"cityId"`
	CityNameEnglish string                 `json:"cityNameEnglish"`
	CityNameArabic  string                 `json:"cityName"`
	X               float64                `json:"x"`
	Y               float64                `json:"y"`
	CreatedAt       string                 `json:"createdAt,omitempty"`
	UpdatedAt       string                 `json:"updatedAt,omitempty"`
	AdditionalInfo  map[string]interface{} `json:"additionalInfo"`
}

func (thisObj *HunterFlowSensor) GetMac() string {
	// return thisObj.MacAddress
	return common.EmptyField
}

func (thisObj *HunterFlowSensor) GetGlobalId() string {
	return thisObj.GlobalId
}

func (thisObj *HunterFlowSensor) GetCommunicationProtocolConfiguration() (string, interface{}) {
	return common.EmptyField, nil
}

func (thisObj *HunterFlowSensor) GetName() string {
	return thisObj.Name
}

func (thisObj *HunterFlowSensor) GetModel() string {
	// return thisObj.Model
	return common.EmptyField
}

func (thisObj *HunterFlowSensor) ValidateModel() error {
	// un implemented until i hava a list of models
	return nil
}

func (thisObj *HunterFlowSensor) GetReferenceName() string {
	return thisObj.GlobalId
}

func (thisObj *HunterFlowSensor) GetClientId() string {
	return ""
}

func (thisObj *HunterFlowSensor) GetIP() string {
	return common.EmptyField
}

func (thisObj *HunterFlowSensor) GetParentAssetId() string {
	return thisObj.AreaId
}

func (thisObj *HunterFlowSensor) GetParentGatewayId() string {
	return thisObj.ControllerId
}

func (thisObj *HunterFlowSensor) GetDeviceType() string {
	return cervello.DeviceTypePeriphral
}

func (thisObj *HunterFlowSensor) GetTags() []string {
	deviceStateTag := common.MockDevice
	if common.IsPhysicalDevice {
		deviceStateTag = common.GisDevice
	}
	return []string{deviceStateTag,
		"Hunter",
		"flowsensor",
		"irrigation",
		thisObj.ControllerName,
		/*thisObj.AreaNameEnglish,*/
		thisObj.AreaNameArabic}
}

func (thisObj *HunterFlowSensor) GetSearchTag() string {
	return "flowsensor"
}

func (thisObj *HunterFlowSensor) SetParentAssetInfo(parentAsset cervello.Asset) error {
	part, err := infrastructuredomain.MigrateCervelloAssetToArea(parentAsset)
	if err != nil {
		return errors.New("error fetching parent area: " + err.Error())
	}

	thisObj.CityId = part.CityId
	thisObj.CityNameEnglish = part.CityNameEnglish
	thisObj.CityNameArabic = part.CityNameArabic
	thisObj.AreaId = part.GlobalId
	thisObj.AreaNameEnglish = part.NameEnglish
	thisObj.AreaNameArabic = part.NameArabic
	return nil
}

func (thisObj *HunterFlowSensor) SetParentGatewayInfo(parentDevice cervello.Device) error {
	controller, err := MigrateHunterControllerFromCervelloDevice(parentDevice)
	if err != nil {
		return errors.New("error fetching parent controller: " + err.Error())
	}
	thisObj.ControllerName = controller.Name
	return nil
}

func (thisObj *HunterFlowSensor) GetLayerType() string {
	return thisObj.LayerType
}

func (thisObj *HunterFlowSensor) Validate() error {
	return crudfunctions.ValidateDeviceEntity(thisObj)
}

func (thisObj *HunterFlowSensor) MigrateFromCsvLine(csvLine []string, keysMap map[string]int) error {
	var err error

	thisObj.GlobalId = csvLine[keysMap["globalId"]]
	thisObj.Name = csvLine[keysMap["name"]]
	thisObj.CityId = csvLine[keysMap["cityId"]]
	thisObj.AreaId = csvLine[keysMap["areaId"]]
	thisObj.X, err = strconv.ParseFloat(csvLine[keysMap["x"]], 64)
	if err != nil {
		return err
	}
	thisObj.Y, err = strconv.ParseFloat(csvLine[keysMap["y"]], 64)
	if err != nil {
		return err
	}
	thisObj.ControllerId = csvLine[keysMap["controllerId"]]
	thisObj.Type = "flowsensor"
	thisObj.LayerType = "point"
	thisObj.AdditionalInfo = common.SetupAdditionalInfo(keysMap, thisObj.GetEssentialKeys(), csvLine)

	return err
}

func (thisObj *HunterFlowSensor) GetEssentialKeys() []string {
	return []string{
		"globalId",
		"integrationId",
		"name",
		"areaId",
		"x",
		"y",
		"controllerId",
	}
}

func (thisObj *HunterFlowSensor) GetNonDuplicatingKeys() []string {
	return []string{
		"globalId",
		"integrationId",
		"name",
	}
}

func (thisObj *HunterFlowSensor) GetParentAssetKey() string {
	return "areaId"
}

func (thisObj *HunterFlowSensor) GetParentGatewayKey() string {
	return "controllerId"
}
