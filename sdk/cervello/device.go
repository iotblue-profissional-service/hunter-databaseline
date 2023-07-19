package cervello

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// DeviceTypeGateWay ...
	DeviceTypeGateWay = "GATEWAY"
	// DeviceTypePeriphral ...
	DeviceTypePeriphral = "PERIPHERAL"
	// DeviceTypeStandalone ...
	DeviceTypeStandalone = "STANDALONE"
	// DeviceProtocolLoraWan ...
	DeviceProtocolLoraWan = "LORAWAN"
	// DeviceProtocolMqtt ...
	DeviceProtocolMqtt = "DEFAULT"
	// DeviceProtocolDefault ...
	DeviceProtocolDefault = "DEFAULT"
	// DeviceProtocolBacnet ...
	DeviceProtocolBacnet = "BACNET"
	// DeviceProtocolModbus ...
	DeviceProtocolModbus = "MODBUS"
)

// Device ...
// cervello device data model
type Device struct {
	ID                    string                 `json:"id,omitempty" validate:"omitempty,uuid"`
	Name                  string                 `json:"name,omitempty"`
	DeviceType            string                 `json:"deviceType,omitempty"`
	Tags                  []string               `json:"tags,omitempty"`
	CommunicationProtocol string                 `json:"communicationProtocol,omitempty"`
	ReferenceName         string                 `json:"referenceName,omitempty"`
	ConnectivityMedia     string                 `json:"connectivityMedia,omitempty"`
	Description           string                 `json:"description,omitempty"`
	AdditionalInfo        string                 `json:"additionalInfo,omitempty"`
	ClientID              string                 `json:"clientId,omitempty"`
	CustomFields          map[string]interface{} `json:"customFields,omitempty"`
	Attributes            map[string]interface{} `json:"attributes,omitempty"`
	MaintenanceMode       bool                   `json:"maintenanceMode,omitempty"`
	LastConnectionStatus  bool                   `json:"lastConnectionStatus,omitempty"`
	ParentGatewayID       string                 `json:"parentGatewayId,omitempty" validate:"omitempty,uuid"`
	ParentGateway         struct {
		ID   string `json:"id,omitempty" validate:"omitempty,uuid"`
		Name string `json:"name,omitempty"`
	} `json:"parentGateway,omitempty"`
	ProtocolConfigurations interface{} `json:"protocolConfigurations,omitempty"`
	CreatedAt              string      `json:"createdAt,omitempty"`
	UpdatedAt              string      `json:"updatedAt,omitempty"`
	LastConnectionTime     string      `json:"lastConnectionTime,omitempty"`
}

// HostPortDeviceProtocolConfiguration .....
type HostPortDeviceProtocolConfiguration struct {
	Host string `json:"host,omitempty" `
	Port int64  `json:"port,omitempty"`
}

// LoraWanDeviceProtocolConfiguration .....
// a lora app , lora gateway , lora profile should be created \
// before creating a lora device with this configuration
type LoraWanDeviceProtocolConfiguration struct {
	Keys struct {
		AppKey string `json:"appKey,omitempty"`
	} `json:"keys,omitempty"`
	Port              int    `json:"port,omitempty"`
	LoraProfileID     string `json:"loraProfileId,omitempty"`
	LoraApplicationID string `json:"loraApplicationId,omitempty"`
}

// BacnetDeviceProtocolConfiguration ...
// bacnet device configuration host (ip address) and port
type BacnetDeviceProtocolConfiguration HostPortDeviceProtocolConfiguration

// ModbusDeviceProtocolConfiguration ...
// modbus device configuration host (ip address) and port
type ModbusDeviceProtocolConfiguration HostPortDeviceProtocolConfiguration

// DeviceAsset ...
// the device data model that resulted from getAssetDevice function
type DeviceAsset struct {
	OrganizationID     string                 `json:"organizationId,omitempty"`
	ApplicationID      string                 `json:"applicationId,omitempty"`
	DeviceID           string                 `json:"deviceId,omitempty"`
	AssetID            string                 `json:"assetId,omitempty"`
	ApplicationName    string                 `json:"applicationName,omitempty"`
	AssetName          string                 `json:"assetName,omitempty"`
	AssetReferenceName string                 `json:"assetReferenceName,omitempty"`
	AssetType          string                 `json:"assetType,omitempty"`
	ReferenceName      string                 `json:"referenceName,omitempty"`
	AdditionalInfo     string                 `json:"additionalInfo,omitempty"`
	CustomFields       map[string]interface{} `json:"customFields,omitempty"`
}

// DeviceTelemetry ...
type DeviceTelemetry struct {
	Data map[string]interface{} `json:"data,omitempty"`
}

// DeviceCredentials ...
type DeviceCredentials struct {
	DeviceID       string `json:"deviceId,omitempty"`
	ClientID       string `json:"clientId,omitempty"`
	AccessKey      string `json:"accessKey,omitempty"`
	AccessToken    string `json:"accessToken,omitempty"`
	OrganizationID string `json:"organizationId,omitempty"`
}

// ModbusConfiguration ...
// these configuration should be added to an existing modbus device
type ModbusConfiguration struct {
	Address       int                          `json:"address,omitempty"`
	Mapping       map[string]map[string]string `json:"mapping,omitempty"`
	SlaveID       int                          `json:"slaveId,omitempty"`
	Quantity      int                          `json:"quantity,omitempty"`
	Sequence      int                          `json:"sequence,omitempty"`
	OperationCode int                          `json:"operationCode,omitempty"`
}

// ModbusConfigurationSchedule ...
// these configuration should be added to an existing modbus device
// it is the period time of pulling data
type ModbusConfigurationSchedule struct {
	Interval int    `json:"interval,omitempty"`
	TimeUnit string `json:"timeUnit,omitempty"`
	Timezone string `json:"timezone,omitempty"`
}

// ModbusDeviceConfig ...
// ModbusConfiguration.Mapping registernumber:{"key":"mykey","type":"TELMETRY"} , type can be ATTRIBUTE (OR) TELMETRY
// ModbusConfiguration.Sequence  = 1 for BIGINDIAN or = 2 for LITTLEINDIAN
// ModbusConfiguration.OperationCode = 3 for ReadingHoldingRegister (OR) 1 for ReadCoils(OR) 2 for ReadDiscreteInput (OR) 4 for ReadInputRegister
// Schedule.TimeUnit = Second | Minute | Hour
// Schedule.TimeZone = Africa/Cairo
type ModbusDeviceConfig struct {
	Configuration []ModbusConfiguration       `json:"configuration,omitempty"`
	Schedule      ModbusConfigurationSchedule `json:"schedule,omitempty"`
}

//bacnetconfiguration start ***********************
// BacnetConfiguration ...
// these configuration should be added to an existing modbus device
type BacnetProperties struct {
	Key        string `json:"key"`
	Type       string `json:"type"`
	DataType   string `json:"dataType"`
	PropertyID int    `json:"propertyId"`
}
type BacnetConfiguration struct {
	ObjectInstance int                `json:"objectInstance"`
	ObjectType     int                `json:"objectType"`
	Peripheral     string             `json:"peripheral"`
	Properties     []BacnetProperties `json:"properties"`
}

// BacnetConfigurationSchedule ...
// these configuration should be added to an existing Bacnet device
// it is the period time of pulling data
type BacnetConfigurationSchedule struct {
	Interval int    `json:"interval,omitempty"`
	TimeUnit string `json:"timeUnit,omitempty"`
	Timezone string `json:"timezone,omitempty"`
}

// ModbusDeviceConfig ...
// BacnetConfiguration.Properties "properties":[{"key":"test","type":"TELEMETRY","dataType":"Number","propertyId":85}],"objectInstance":1}] , type can be ATTRIBUTE (OR) TELMETRY
// ModbusConfiguration.Sequence  = 1 for BIGINDIAN or = 2 for LITTLEINDIAN
// ModbusConfiguration.OperationCode = 3 for ReadingHoldingRegister (OR) 1 for ReadCoils(OR) 2 for ReadDiscreteInput (OR) 4 for ReadInputRegister
// Schedule.TimeUnit = Second | Minute | Hour
// Schedule.TimeZone = Africa/Cairo
type BacnetDeviceConfig struct {
	ID            string `json:"id,omitempty"`
	Configuration struct {
		Object []BacnetConfiguration `json:"objects"`
	} `json:"configuration,omitempty"`
	Schedule BacnetConfigurationSchedule `json:"schedule,omitempty"`
}

// LoraActivationParams ....

type LoraActivationParams struct {
	ClientID       string `json:"clientId,omitempty"`
	ApplicationKey string `json:"appKey,omitempty"`
	DeviceID       string
}

type LoraActivationParamsABP struct {
	ClientID   string               `json:"clientId,omitempty"`
	Activation LoraActivationObject `json:"activation"`
}
type LoraActivationObject struct {
	ApplicationSessionKey string `json:"appSKey"`
	NetworkSessionKey1    string `json:"fNwkSIntKey"`
	NetworkSessionKey2    string `json:"nwkSEncKey"`
	NetworkSessionKey3    string `json:"sNwkSIntKey"`
	UplinkFrameCounter    int    `json:"fCntUp"`
	DownLinkFrameCounter  int    `json:"nFCntDown"`
	DeviceAddress         string `json:"devAddr"`
}

// GetDevice ...
// using the device refrence name
func GetDevice(referenceID string, privateToken string, configs ...CervelloConfigurations) (*Device, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}
	// GET https://api.staging.cervello.io/application/v1/organizations/0feb9329-cf80-4c61-812b-c67eb74d5c33/devices
	// ?pageSize=1&filters[0][key]=referenceName&filters[0][operator]=eq&filters[0][value]=2388sa
	// headers [Authorization: "Bearer {token}"]
	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/devices",
		config.OrganizationID,
	)

	var responseBodyObject *getDeviceResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource: resource,
		MetaData: "GetDevice",
		Token:    privateToken,
		QueryParams: QueryParams{
			Filters:       []Filter{{Key: "referenceName", Op: "eq", Value: referenceID}},
			PaginationObj: Pagination{PageNumber: 1, PageSize: 1},
		},
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetDevice unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "unexpected error", string(response.ResponseBody))
	}

	if responseBodyObject.Result.Devices == nil || len(responseBodyObject.Result.Devices) == 0 {
		return &Device{}, nil
	}

	return &responseBodyObject.Result.Devices[0], nil
}

// GetDeviceByID ...
// using the device id
func GetDeviceByID(id string, privateToken string, configs ...CervelloConfigurations) (*Device, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}
	//https://api.staging.cervello.io/device/v1/organizations/8bd3d7dc-1b71-4ae7-9d23-213f7e83a102/devices/a2179afd-1776-469d-814b-9ab877fde5c4
	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/devices/%s",
		config.OrganizationID,
		id,
	)

	var responseBodyObject *getDeviceBYIDResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource: resource,
		MetaData: "GetDeviceByID",
		Token:    privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetDeviceByID unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "unexpected error", string(response.ResponseBody))
	}

	return &responseBodyObject.Result, nil

}

// GetDeviceFiltered ...
//
// to get device by refrence id set filter --- > key = referenceName , op=eq , value = theRefrencceNAmeValue
func GetDeviceFiltered(queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) (*Device, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}
	// GET https://api.staging.cervello.io/application/v1/organizations/0feb9329-cf80-4c61-812b-c67eb74d5c33/devices
	// ?pageSize=1&filters[0][key]=referenceName&filters[0][operator]=eq&filters[0][value]=2388sa
	// headers [Authorization: "Bearer {token}"]

	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/devices",
		config.OrganizationID,
	)

	var responseBodyObject *getDeviceResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource:    resource,
		MetaData:    "GetDeviceByID",
		Token:       privateToken,
		QueryParams: queryParam,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetDeviceByID unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "unexpected error", string(response.ResponseBody))
	}
	if responseBodyObject.Result.Devices == nil || len(responseBodyObject.Result.Devices) == 0 {
		return &Device{}, nil
	}

	return &responseBodyObject.Result.Devices[0], nil

}

// GetDeviceAttributes ...
func GetDeviceAttributes(deviceID string, privateToken string, configs ...CervelloConfigurations) (map[string]interface{}, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}

	resource := fmt.Sprintf(
		"/data/v1/organizations/%s/devices/%s/attributes",
		config.OrganizationID,
		deviceID,
	)

	attributes := new(map[string]interface{})

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource:    resource,
		MetaData:    "GetDeviceByID",
		Token:       privateToken,
		QueryParams: QueryParams{PaginationObj: Pagination{PageNumber: 1, PageSize: 10000000}},
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &attributes); err != nil {
		internalLog("error", "GetDeviceByID unmarshal response", err)
		return nil, err
	}

	return *attributes, nil
}

// CreateDevice .....
func CreateDevice(device Device, privateToken string, configs ...CervelloConfigurations) (*Device, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}

	resource := fmt.Sprintf(
		"/compose/v1/organizations/%s/devices",
		config.OrganizationID,
	)

	var responseBodyObject postDeviceResponse

	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "CreateDevice",
		Token:    privateToken,
		Body:     device,
	})

	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "CreateDevice unmarshal response", err)
		return nil, err
	}
	if responseBodyObject.Code != 2001 {
		return nil, errors.New("Failed to CreateDevice: " + string(response.ResponseBody))
	}
	return &responseBodyObject.Result.Device, nil

}

// UpdateDevice .....
func UpdateDevice(deviceID string, device Device, privateToken string, configs ...CervelloConfigurations) error {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return err
	}

	resource := fmt.Sprintf(
		"/compose/v1/organizations/%s/devices/%s",
		config.OrganizationID,
		deviceID,
	)

	var responseBodyObject baseResponse

	response, err := makeHTTPRequest("PUT", htppRequest{
		Resource: resource,
		MetaData: "UpdateDevice",
		Token:    privateToken,
		Body:     device,
	})

	if err != nil {
		internalLog("error", err)
		return err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "UpdateDevice unmarshal response", err)
		return err
	}
	if responseBodyObject.Code != 2002 {
		return errors.New("Failed to UpdateDevice: " + string(response.ResponseBody))
	}
	return nil

}

// AssignDeviceToApplication .....
func AssignDeviceToApplication(deviceID string, privateToken string, configs ...CervelloConfigurations) (bool, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return false, err
	}
	// POST https://api.staging.cervello.io/compose/v1/organizations/dbb02a67-6ab6-4485-b955-e485ade0a003/devices/055685c4-06dc-4996-b36c-7989c354ef66/applications
	// {"applications":["7d12790c-5bbd-4168-9689-34436ebd8543"]}
	// headers [Authorization]
	//http://api.cervello.local/compose/v1/organizations/3616505e-0a6d-4893-98f9-18c176775e15/devices/9c7191ec-a8c0-4ba4-9ebe-43a5d9165ae6/application
	resource := "/compose/v1/organizations/" + config.OrganizationID + "/devices/" + deviceID + "/applications"

	var applicationsIDs = map[string]interface{}{
		"applications": []string{
			envApplicationID,
		},
	}

	// log.Printf(applicationsIDs)

	var responseBodyObject baseResponse

	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "AssignDeviceToApplication",
		Token:    privateToken,
		Body:     applicationsIDs,
	})

	if err != nil {
		internalLog("error", err)
		return false, err
	}
	// log.Printf(string(response.ResponseBody))
	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "AssignDeviceToApplication unmarshal response", err)
		return false, err
	}
	if responseBodyObject.Code != 2001 {
		return false, errors.New("Failed to AssignDeviceToApplication: " + string(response.ResponseBody))
	}
	return true, nil

}

// CreateDeviceAttributes .....
func CreateDeviceAttributes(deviceID string, attributes map[string]interface{}, privateToken string, configs ...CervelloConfigurations) (bool, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return false, err
	}

	resource := "/data/v1/organizations/" + config.OrganizationID + "/devices/" + deviceID + "/attributes"

	var attributeBody = map[string]interface{}{
		"data": attributes,
	}

	var responseBodyObject baseResponse

	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "CreateDeviceAttributes",
		Token:    privateToken,
		Body:     attributeBody,
	})

	if err != nil {
		internalLog("error", err)
		return false, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "CreateDeviceAttributes unmarshal response", err)
		return false, err
	}
	if responseBodyObject.Code != 2001 {
		return false, errors.New("Failed to CreateDeviceAttributes: " + string(response.ResponseBody))
	}

	return true, nil
}

// GetDeviceAssetIDs .....
func GetDeviceAssetIDs(deviceID string, privateToken string, configs ...CervelloConfigurations) ([]string, error) {
	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return nil, err
	}
	// GET https://api.staging.cervello.io/application/v1/organizations/ac3b5721-33ee-4e19-b3a8-60322e5044c7/applications/85430088-b300-4fc0-94e1-9df01b9db32f/devices/69481928-ef39-4d07-9a37-4b667882c3ea/assets
	var ids []string
	resource := fmt.Sprintf(
		"/application/v1/organizations/%s/applications/%s/devices/%s/assets",
		config.OrganizationID,
		config.ApplicationID,
		deviceID,
	)

	var responseBodyObject *getDeviceAssetsResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource: resource,
		MetaData: "GetDeviceAssetsNames",
		Token:    privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetDeviceAssetsNames unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "unexpected error", string(response.ResponseBody))
	}

	for _, deviceAsset := range responseBodyObject.Result.DeviceAssets {
		ids = append(ids, deviceAsset.AssetID)
	}
	return ids, nil
}

// GetDeviceLastTelmetries .... key must be empty to get All keys
func GetDeviceLastTelmetries(deviceID string, key string, queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) ([]Telmetry, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}
	//https://api.release.cervello.io/data/v1/organizations/724a3eb5-6a76-4fb6-9a6d-499491e9c1d7/devices/6e4600fc-8f84-494b-a335-d42af311fb18/telemetries/latest?pageNumber=1&pageSize=10&aggregation=last
	resource := fmt.Sprintf(
		"/data/v1/organizations/%s/devices/%s/telemetries",
		config.OrganizationID,
		deviceID,
	)

	data := url.Values{}
	data = addQueryParamsToRequest(data, queryParam)

	switch key {
	case "":
	default:
		data.Set("filters[0][key]", "key")
		data.Set("filters[0][operator]", "eq")
		data.Set("filters[0][value]", key)
	}

	data.Set("aggregation", "last")
	u, _ := url.ParseRequestURI(envAPIURL)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := u.String() // "https://api.com/user/"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", urlStr, nil) // URL-encoded payload
	if err != nil {
		internalLog("error", "GetDeviceLastTelmetries-NewRequest: ")
		internalLog("error", err)
		return []Telmetry{}, err
	}

	addAuthHeader(req, privateToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		internalLog("error", "GetDeviceLastTelmetries-Do: ")
		internalLog("error", err)
		return []Telmetry{}, err
	}

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		internalLog("error", "GetDeviceLastTelmetries-ReadAll: ")
		internalLog("error", err)
		return []Telmetry{}, err
	}

	resp.Body.Close()

	// Assets := new(getAssetResponse)
	var tagTelemetries getTagTelemtriesResponse

	err = json.Unmarshal(f, &tagTelemetries)
	if err != nil {
		internalLog("error", "GetDeviceLastTelmetries-Unmarshal: ")
		internalLog("error", err)
		return []Telmetry{}, err
	}

	return tagTelemetries.Result.Telmetries, nil
}

// DeleteDevice .....
// if gateway device then all children preiphrals will be deleted
func DeleteDevice(deviceID string, privateToken string, configs ...CervelloConfigurations) error {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return err
	}
	// https://api.release.cervello.io/device/v1/organizations/ffae83db-f249-4530-8db3-7298fdb75777/devices/ad56b53e-792b-480f-adbe-38843ef26a31
	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/devices/%s",
		config.OrganizationID,
		deviceID,
	)

	var responseBodyObject baseResponse

	response, err := makeHTTPRequest("DELETE", htppRequest{
		Resource:    resource,
		MetaData:    "DeleteDevice",
		Token:       privateToken,
		QueryParams: QueryParams{Custom: map[string]interface{}{"force": "true"}},
	})
	if err != nil {
		internalLog("error", err)
		return err
	}
	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "DeleteDevice unmarshal response", err)
		return err
	}
	if responseBodyObject.Code != 2004 {
		return errors.New("Failed to DeleteDevice " + string(response.ResponseBody))
	}
	return nil

}

// GetDeviceAssetsNames .....
func GetDeviceAssetsNames(deviceID string, privateToken string, configs ...CervelloConfigurations) ([]string, error) {
	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return nil, err
	}
	// GET https://api.staging.cervello.io/application/v1/organizations/ac3b5721-33ee-4e19-b3a8-60322e5044c7/applications/85430088-b300-4fc0-94e1-9df01b9db32f/devices/69481928-ef39-4d07-9a37-4b667882c3ea/assets
	var names []string
	resource := fmt.Sprintf(
		"/application/v1/organizations/%s/applications/%s/devices/%s/assets",
		config.OrganizationID,
		config.ApplicationID,
		deviceID,
	)

	var responseBodyObject *getDeviceAssetsResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource: resource,
		MetaData: "GetDeviceAssetsNames",
		Token:    privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetDeviceAssetsNames unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "unexpected error", string(response.ResponseBody))
	}

	for _, deviceAsset := range responseBodyObject.Result.DeviceAssets {
		names = append(names, deviceAsset.AssetName)
	}
	return names, nil
}

// GetDeviceAssets .....
func GetDeviceAssets(deviceID string, privateToken string, configs ...CervelloConfigurations) ([]DeviceAsset, error) {
	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return nil, err
	}
	// GET https://api.staging.cervello.io/application/v1/organizations/ac3b5721-33ee-4e19-b3a8-60322e5044c7/applications/85430088-b300-4fc0-94e1-9df01b9db32f/devices/69481928-ef39-4d07-9a37-4b667882c3ea/assets
	resource := fmt.Sprintf(
		"/application/v1/organizations/%s/applications/%s/devices/%s/assets",
		config.OrganizationID,
		config.ApplicationID,
		deviceID,
	)

	var responseBodyObject *getDeviceAssetsResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource:    resource,
		MetaData:    "GetDeviceAssets",
		Token:       privateToken,
		QueryParams: QueryParams{Custom: map[string]interface{}{"ancestor": "1"}},
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetDeviceAssets unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "unexpected error", string(response.ResponseBody))
	}

	return responseBodyObject.Result.DeviceAssets, nil

}

// ChangeDeviceMaintananceMode .....
func ChangeDeviceMaintananceMode(deviceIDs []string, isMaintanance bool, privateToken string, configs ...CervelloConfigurations) error {
	//device/v1/organizations/9a450c80-8511-4de5-8a1c-a3ecc6e9c862/devices/off-maintenance
	// device/v1/organizations/9a450c80-8511-4de5-8a1c-a3ecc6e9c862/devices/on-maintenance

	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return err
	}
	// let initial resource removing from maintanance resource
	resource := "/device/v1/organizations/" + config.OrganizationID + "/devices/off-maintenance"

	if isMaintanance {
		resource = "/device/v1/organizations/" + config.OrganizationID + "/devices/on-maintenance"
	}

	deviceIDsStruct := struct {
		DeviceIds []string `json:"devicesIds"`
	}{
		DeviceIds: deviceIDs,
	}

	var responseBodyObject baseResponse

	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "ChangeDeviceMaintananceMode",
		Token:    privateToken,
		Body:     deviceIDsStruct,
	})

	if err != nil {
		internalLog("error", err)
		return err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "ChangeDeviceMaintananceMode unmarshal response", err)
		return err
	}
	if responseBodyObject.Code != 2002 {
		return errors.New("Failed to ChangeDeviceMaintananceMode: " + string(response.ResponseBody))
	}

	return nil
}

// GetGateWayPheriphrals .....
func GetGateWayPheriphrals(gateWayID string, queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) ([]Device, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}
	//https://api.release.cervello.io/device/v1/organizations/e4550d32-48f7-4f9c-bdf9-95df8f788140/gateways/b744e761-e5f8-40af-86d9-e16d31dac269/peripherals
	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/gateways/%s/peripherals",
		config.OrganizationID,
		gateWayID,
	)

	var responseBodyObject *getDeviceResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource:    resource,
		MetaData:    "GetGateWayPheriphrals",
		Token:       privateToken,
		QueryParams: queryParam,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetGateWayPheriphrals unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "unexpected error", string(response.ResponseBody))
	}
	if responseBodyObject.Result.Devices == nil || len(responseBodyObject.Result.Devices) == 0 {
		return []Device{}, nil
	}

	return responseBodyObject.Result.Devices, nil
}

// PublishDeviceTelemetry .....
func PublishDeviceTelemetry(credentials DeviceCredentials, telemetry DeviceTelemetry) error {
	data := url.Values{}
	data.Set("t", fmt.Sprintf(
		"t=/device/%s/telemetry",
		credentials.ClientID,
	))

	deviceJSON, err := json.Marshal(telemetry)
	if err != nil {
		internalLog("error", err)
		return err
	}

	u, _ := url.ParseRequestURI(envBrokerAPIURL)
	u.RawQuery = data.Get("t")
	urlStr := u.String()

	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(deviceJSON)) // URL-encoded payload
	if err != nil {
		internalLog("error", err)
		return err
	}

	auth := credentials.AccessKey + ":" + credentials.AccessToken
	basicToken := base64.StdEncoding.EncodeToString([]byte(auth))

	req.Header.Add("Authorization", "Basic "+basicToken)
	req.Header.Add("x-client-id", credentials.ClientID)

	req.Header.Add("Content-Type", "application/json")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		internalLog("error", "hello", err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		internalLog("error", err)
		return err
	}

	var response baseResponse
	err = json.Unmarshal(body, &response)

	if response.Code != 2005 {
		return errors.New("Failed to publish device telemetry .Error: " + response.MessageKey)
	}

	return nil
}

// GetDeviceCredentials .....
func GetDeviceCredentials(deviceID string, privateToken string, configs ...CervelloConfigurations) (*DeviceCredentials, error) {

	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}

	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/devices/%s/credentials",
		config.OrganizationID,
		deviceID,
	)

	var responseBodyObject deviceCredentialsResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource: resource,
		MetaData: "GetDeviceCredentials",
		Token:    privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetDeviceCredentials unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "unexpected error", string(response.ResponseBody))
	}

	return &responseBodyObject.Result, nil

}

// GenerateDeviceCredentials .....
func GenerateDeviceCredentials(deviceID string, privateToken string, configs ...CervelloConfigurations) (*DeviceCredentials, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}

	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/devices/%s/credentials",
		config.OrganizationID,
		deviceID,
	)

	var responseBodyObject deviceCredentialsResponse

	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "GenerateDeviceCredentials",
		Token:    privateToken,
	})

	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GenerateDeviceCredentials unmarshal response", err)
		return nil, err
	}
	if responseBodyObject.Code != 2003 {
		return nil, errors.New("Failed to GenerateDeviceCredentials: " + string(response.ResponseBody))
	}

	return &responseBodyObject.Result, nil
}

// GetOrgDevicesFiltered ...
//
// example Get connnected devices
//
//	filters.key = lastConnectionStatus
// 	filters.Op = eq
//  filters.Value = true
func GetOrgDevicesFiltered(queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) ([]Device, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}

	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/devices",
		config.OrganizationID,
	)

	var responseBodyObject getDeviceResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource:    resource,
		QueryParams: queryParam,
		MetaData:    "GetOrgDevicesFiltered",
		Token:       privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetOrgDevicesFiltered unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "unexpected error", string(response.ResponseBody))
	}

	return responseBodyObject.Result.Devices, nil
}

// GetModbusDeviceConfig ....
func GetModbusDeviceConfig(deviceID string, privateToken string, configs ...CervelloConfigurations) (*ModbusDeviceConfig, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}
	//Request URL: https://api.staging.cervello.io/device/v1/organizations/8bd3d7dc-1b71-4ae7-9d23-213f7e83a102/devices/71ff2411-ff11-46cc-80f7-9cba4e527517/modbus
	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/devices/%s/modbus",
		config.OrganizationID,
		deviceID,
	)

	var responseBodyObject *getModbusConfigurationResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource: resource,
		MetaData: "GetModbusDeviceConfig",
		Token:    privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetModbusDeviceConfig unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "unexpected error", string(response.ResponseBody))
	}

	return &ModbusDeviceConfig{
		Configuration: responseBodyObject.Result.Configuration,
		Schedule:      responseBodyObject.Result.Schedule,
	}, nil

}

//UpdateModbusDeviceConfig ....
// note this functions overrides the old configs and schedule not adding new configs to the odld ones
func UpdateModbusDeviceConfig(deviceID string, modbusConfig *ModbusDeviceConfig, privateToken string, configs ...CervelloConfigurations) error {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return err
	}
	// https://api.staging.cervello.io/device/v1/organizations/8bd3d7dc-1b71-4ae7-9d23-213f7e83a102/devices/71ff2411-ff11-46cc-80f7-9cba4e527517/modbus
	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/devices/%s/modbus",
		config.OrganizationID,
		deviceID,
	)

	var responseBodyObject baseResponse
	response, err := makeHTTPRequest("PUT", htppRequest{
		Resource: resource,
		Body:     modbusConfig,
		MetaData: "UpdateModbusDeviceConfig",
		Token:    privateToken,
	})

	if err != nil {
		internalLog("error", err)
		return err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "UpdateModbusDeviceConfig unmarshal response", err)
		return err
	}

	if responseBodyObject.Code != 2002 {
		return fmt.Errorf("%v-%v", "unexpected error while UpdateModbusDeviceConfig", string(response.ResponseBody))
	}

	return nil
}

// GetBacnetDeviceConfig ....
func GetBacnetDeviceConfig(deviceID string, privateToken string, configs ...CervelloConfigurations) (*BacnetDeviceConfig, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}
	//Request URL: https://api.staging.cervello.io/device/v1/organizations/8bd3d7dc-1b71-4ae7-9d23-213f7e83a102/devices/71ff2411-ff11-46cc-80f7-9cba4e527517/modbus
	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/devices/%s/bacnet",
		config.OrganizationID,
		deviceID,
	)

	var responseBodyObject *getBacnetConfigurationResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource: resource,
		MetaData: "GetBacnetDeviceConfig",
		Token:    privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetBacnetDeviceConfig unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "unexpected error", string(response.ResponseBody))
	}

	return &BacnetDeviceConfig{
		ID: responseBodyObject.Result.ID,
		Configuration: struct {
			Object []BacnetConfiguration "json:\"objects\""
		}{Object: responseBodyObject.Result.Configuration.Objects},
		Schedule: responseBodyObject.Result.Schedule,
	}, nil

}

//UpdateBacnetDeviceConfig ....
// note this functions overrides the old configs and schedule not adding new configs to the odld ones
func CreateBacnetDeviceConfig(deviceID string, config *BacnetDeviceConfig, privateToken string) error {
	// https://api.cervello.local/device/v1/organizations/6d992053-df15-4610-9e2d-8a17c034cc9f/devices/0fab3e08-8b31-4a04-ba96-ed596bad7671/bacnet
	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/devices/%s/bacnet",
		envOrganizationID,
		deviceID,
	)

	var responseBodyObject baseResponse
	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		Body:     config,
		MetaData: "UpdateBacnetDeviceConfig",
		Token:    privateToken,
	})

	if err != nil {
		internalLog("error", err)
		return err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "UpdateBacnetDeviceConfig unmarshal response", err)
		return err
	}

	if responseBodyObject.Code != 2001 {
		return fmt.Errorf("%v-%v", "unexpected error while UpdateBacnetDeviceConfig", string(response.ResponseBody))
	}

	return nil
}

//UpdateBacnetDeviceConfig ....
// note this functions overrides the old configs and schedule not adding new configs to the odld ones
func UpdateBacnetDeviceConfig(deviceID string, bacnetConfig *BacnetDeviceConfig, privateToken string, configs ...CervelloConfigurations) error {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return err
	}

	// https://api.cervello.local/device/v1/organizations/6d992053-df15-4610-9e2d-8a17c034cc9f/devices/0fab3e08-8b31-4a04-ba96-ed596bad7671/bacnet
	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/devices/%s/bacnet/%s",
		config.OrganizationID,
		deviceID,
		bacnetConfig.ID,
	)

	var responseBodyObject baseResponse
	response, err := makeHTTPRequest("PUT", htppRequest{
		Resource: resource,
		Body:     bacnetConfig,
		MetaData: "UpdateBacnetDeviceConfig",
		Token:    privateToken,
	})

	if err != nil {
		internalLog("error", err)
		return err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "UpdateModbusDeviceConfig unmarshal response", err)
		return err
	}

	if responseBodyObject.Code != 2002 {
		return fmt.Errorf("%v-%v", "unexpected error while UpdateModbusDeviceConfig", string(response.ResponseBody))
	}

	return nil
}

// SetLoraDeviceApplicationKey ....

func SetLoraDeviceApplicationKey(loraActivationParams LoraActivationParams, privateToken string, configs ...CervelloConfigurations) error {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return err
	}
	//Request URL: http://api.cervello.local/device/v1/lora/organizations/3616505e-0a6d-4893-98f9-18c176775e15/devices/61ee8904-42d6-4239-b1dc-7f487c9f64ce/keys

	// POST

	// body {clientId: "b6886dae4e13bfa7", appKey: "6435835d11d19e9382b6361ff352434d"}
	resource := "/device/v1/lora/organizations/" + config.OrganizationID + "/devices/" + loraActivationParams.DeviceID + "/keys"

	var body = map[string]interface{}{
		"clientId": loraActivationParams.ClientID,
		"appKey":   loraActivationParams.ApplicationKey,
	}
	var responseBodyObject baseResponse
	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "SetLoraDeviceApplicationKey",
		Token:    privateToken,
		Body:     body,
	})
	if err != nil {
		internalLog("error", err)
		return err
	}
	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "SetLoraDeviceApplicationKey unmarshal response", err)
		return err
	}

	if responseBodyObject.MessageKeys != "success" {
		return errors.New("Failed to SetLoraDeviceApplicationKey: " + string(response.ResponseBody))
	}
	return nil

}

func ActivateLoraDeviceABP(loraActivationParams LoraActivationParamsABP, deviceID string, privateToken string, configs ...CervelloConfigurations) error {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return err
	}
	//Request URL: http://api.cervello.local/device/v1/lora/organizations/3616505e-0a6d-4893-98f9-18c176775e15/devices/61ee8904-42d6-4239-b1dc-7f487c9f64ce/keys

	// POST

	// body {clientId: "b6886dae4e13bfa7", appKey: "6435835d11d19e9382b6361ff352434d"}
	resource := "/device/v1/lora/organizations/" + config.OrganizationID + "/devices/" + deviceID + "/lora"

	var body = loraActivationParams
	var responseBodyObject baseResponse
	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "ActivateLoraDeviceABP",
		Token:    privateToken,
		Body:     body,
	})
	if err != nil {
		internalLog("error", err)
		return err
	}
	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "ActivateLoraDeviceABP unmarshal response", err)
		return err
	}

	if responseBodyObject.MessageKeys != "success" {
		return errors.New("Failed to ActivateLoraDeviceABP: " + string(response.ResponseBody))
	}
	return nil

}
