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

type HunterFlowZone struct {
	GlobalId        string                 `json:"globalId"`
	Name            string                 `json:"name"`
	Type            string                 `json:"type"`
	LayerType       string                 `json:"layerType"`
	ControllerName  string                 `json:"controllerName"`
	ControllerId    string                 `json:"controllerId"`
	Order           int64                  `json:"order"`
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

func (thisObj *HunterFlowZone) GetMac() string {
	// return thisObj.MacAddress
	return common.EmptyField
}

func (thisObj *HunterFlowZone) GetGlobalId() string {
	return thisObj.GlobalId
}

func (thisObj *HunterFlowZone) GetCommunicationProtocolConfiguration() (string, interface{}) {
	return common.EmptyField, nil
}

func (thisObj *HunterFlowZone) GetName() string {
	return thisObj.Name
}

func (thisObj *HunterFlowZone) GetModel() string {
	// return thisObj.Model
	return common.EmptyField
}

func (thisObj *HunterFlowZone) ValidateModel() error {
	// un implemented until i hava a list of models
	return nil
}

func (thisObj *HunterFlowZone) GetReferenceName() string {
	return thisObj.GlobalId
}

func (thisObj *HunterFlowZone) GetClientId() string {
	return ""
}

func (thisObj *HunterFlowZone) GetIP() string {
	return common.EmptyField
}

func (thisObj *HunterFlowZone) GetParentAssetId() string {
	return thisObj.AreaId
}

func (thisObj *HunterFlowZone) GetParentGatewayId() string {
	return thisObj.ControllerId
}

func (thisObj *HunterFlowZone) GetDeviceType() string {
	return cervello.DeviceTypePeriphral
}

func (thisObj *HunterFlowZone) GetTags() []string {
	deviceStateTag := common.MockDevice
	if common.IsPhysicalDevice {
		deviceStateTag = common.GisDevice
	}
	return []string{deviceStateTag,
		"Hunter",
		"flowsensor",
		"flowzone",
		"irrigation",
		thisObj.ControllerName,
		/*thisObj.AreaNameEnglish,*/
		thisObj.AreaNameArabic}
}

func (thisObj *HunterFlowZone) GetSearchTag() string {
	return "flowzone"
}

func (thisObj *HunterFlowZone) SetParentAssetInfo(parentAsset cervello.Asset) error {
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

func (thisObj *HunterFlowZone) SetParentGatewayInfo(parentDevice cervello.Device) error {
	controller, err := MigrateHunterControllerFromCervelloDevice(parentDevice)
	if err != nil {
		return errors.New("error fetching parent controller: " + err.Error())
	}
	thisObj.ControllerName = controller.Name
	return nil
}

func (thisObj *HunterFlowZone) GetLayerType() string {
	return thisObj.LayerType
}

func (thisObj *HunterFlowZone) Validate() error {
	return crudfunctions.ValidateDeviceEntity(thisObj)
}

func (thisObj *HunterFlowZone) MigrateFromCsvLine(csvLine []string, keysMap map[string]int) error {
	var err error

	thisObj.GlobalId = csvLine[keysMap["globalId"]]
	thisObj.Name = fmt.Sprintf("%s-%s", csvLine[keysMap["name"]], csvLine[keysMap["controllerId"]])
	thisObj.CityId = csvLine[keysMap["cityId"]]
	thisObj.AreaId = csvLine[keysMap["areaId"]]
	thisObj.Order, err = strconv.ParseInt(csvLine[keysMap["order"]], 10, 64)
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
	thisObj.ControllerId = csvLine[keysMap["controllerId"]]
	thisObj.Type = csvLine[keysMap["type"]]
	thisObj.LayerType = "point"
	thisObj.AdditionalInfo = common.SetupAdditionalInfo(keysMap, thisObj.GetEssentialKeys(), csvLine)

	return err
}

func (thisObj *HunterFlowZone) GetEssentialKeys() []string {
	return []string{
		"globalId",
		"integrationId",
		"name",
		"areaId",
		"x",
		"y",
		"controllerId",
		"order",
	}
}

func (thisObj *HunterFlowZone) GetNonDuplicatingKeys() []string {
	return []string{
		"globalId",
		"integrationId",
		"name",
	}
}

func (thisObj *HunterFlowZone) GetParentAssetKey() string {
	return "areaId"
}

func (thisObj *HunterFlowZone) GetParentGatewayKey() string {
	return "controllerId"
}
