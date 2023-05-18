package cervello

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	// LoraService ...
	LoraService loraServiceInterface = &loraService{}
)

type loraService struct {
}

type MultiCastGroupCommand struct {
	Data  string `json:"data"`
	Fport int    `json:"fPort"`
	Fcnt  int    `json:"fCnt"`
}

type LoraGateway struct {
	ID                string      `json:"id,omitempty" validate:"omitempty,uuid"`
	Name              string      `json:"name,omitempty"`
	Description       string      `json:"description,omitempty"`
	OrganizationId    string      `json:"organizationID,omitempty"`
	NetworkServerId   string      `json:"networkServerID,omitempty"`
	NetworkServerName string      `json:"networkServerName,omitempty"`
	CreatedAt         string      `json:"createdAt,omitempty"`
	UpdatedAt         string      `json:"updatedAt,omitempty"`
	Location          LocationObj `json:"location,omitempty"`
}

type getLoraGatewayByIdResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys,omitempty"`
	MessageKey  string `json:"messageKey,omitempty"`
	Message     string `json:"message,omitempty"`
	Result      struct {
		TotalCount string        `json:"totalCount"`
		Data       []LoraGateway `json:"data"`
	} `json:"result"`
}

type LocationObj struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Altitude  float64 `json:"altitude,omitempty"`
	Source    string  `json:"source,omitempty"`
	Accuracy  float64 `json:"accuracy,omitempty"`
}

type postLoraGatewayResponse struct {
	Code        int                    `json:"code,omitempty"`
	MessageKeys string                 `json:"messageKeys,omitempty"`
	MessageKey  string                 `json:"messageKey,omitempty"`
	Message     string                 `json:"message,omitempty"`
	Result      map[string]interface{} `json:"result,omitempty"`
}

type loraServiceInterface interface {
	QueueMultiCastGroupCommand(multiCastGroupID string, command MultiCastGroupCommand, privateToken string, configs ...CervelloConfigurations) error
	AddDeviceToMultiCastGroup(multiCastGroupID string, deviceId string, deviceEUI string, privateToken string, configs ...CervelloConfigurations) error
	RemoveDeviceFromMultiCastGroup(multiCastGroupID string, deviceId string, deviceEUI string, privateToken string, configs ...CervelloConfigurations) error
	GetMultiCastGroupDevices(multiCastGroupID string, privateToken string, configs ...CervelloConfigurations) ([]Device, error)
	CreateGateway(gateway LoraGateway, privateToken string, configs ...CervelloConfigurations) (*map[string]interface{}, error)
	UpdateGateway(id string, gateway LoraGateway, privateToken string, configs ...CervelloConfigurations) (*map[string]interface{}, error)
	DeleteGateway(id string, privateToken string, configs ...CervelloConfigurations) (*map[string]interface{}, error)
	GetGatewayByID(ID string, privateToken string, configs ...CervelloConfigurations) (*LoraGateway, error)
	GetGateways(privateToken string, configs ...CervelloConfigurations) ([]LoraGateway, error)
}

func (lora *loraService) QueueMultiCastGroupCommand(multiCastGroupID string, command MultiCastGroupCommand, privateToken string, configs ...CervelloConfigurations) error {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return err
	}

	//http://localhost:3001/device/v1/lora/organizations/cbb85128-ff08-44d5-a568-972d4bd151d2/multicastGroups/4741ae7f-5769-493d-b5b7-e67523c40a93/queue

	resource := "/device/v1/lora/organizations/" + config.OrganizationID + "/multicastGroups/" + multiCastGroupID + "/queue"

	var responseBodyObject baseResponse
	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "QueueMultiCastGroupCommand",
		Token:    privateToken,
		Body:     command,
	})
	if err != nil {
		internalLog("error", err)
		return err
	}
	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "QueueMultiCastGroupCommand unmarshal response", err)
		return err
	}

	if responseBodyObject.MessageKeys != "success" {
		return errors.New("Failed to QueueMultiCastGroupCommand: " + string(response.ResponseBody))
	}
	return nil
}

func (lora *loraService) AddDeviceToMultiCastGroup(multiCastGroupID string, deviceId string, deviceEUI string, privateToken string, configs ...CervelloConfigurations) error {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return err
	}
	//http://localhost:3001/device/v1/lora/organizations/cbb85128-ff08-44d5-a568-972d4bd151d2/multicastGroups/4741ae7f-5769-493d-b5b7-e67523c40a93/devices

	resource := "/device/v1/lora/organizations/" + config.OrganizationID + "/multicastGroups/" + multiCastGroupID + "/devices"

	var body = map[string]interface{}{
		"devEUI":   deviceEUI,
		"deviceId": deviceId,
	}
	var responseBodyObject baseResponse
	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "AddDeviceToMultiCastGroup",
		Token:    privateToken,
		Body:     body,
	})
	if err != nil {
		internalLog("error", err)
		return err
	}
	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "AddDeviceToMultiCastGroup unmarshal response", err)
		return err
	}

	if responseBodyObject.MessageKeys != "success" {
		return errors.New("Failed to AddDeviceToMultiCastGroup: " + string(response.ResponseBody))
	}
	return nil
}

func (lora *loraService) RemoveDeviceFromMultiCastGroup(multiCastGroupID string, deviceId string, deviceEUI string, privateToken string, configs ...CervelloConfigurations) error {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return err
	}
	//http://api.cervello.local/device/v1/lora/organizations/:organizationId/multicastGroups/:multicastGroupId/devices/:deviceId/loraDevice/:devEUI
	resource := "/device/v1/lora/organizations/" + config.OrganizationID + "/multicastGroups/" + multiCastGroupID + "/devices/" + deviceId + "/loraDevice/" + deviceEUI

	var responseBodyObject baseResponse
	response, err := makeHTTPRequest("DELETE", htppRequest{
		Resource: resource,
		MetaData: "RemoveDeviceToMultiCastGroup",
		Token:    privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return err
	}
	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "RemoveDeviceToMultiCastGroup unmarshal response", err)
		return err
	}

	if responseBodyObject.MessageKeys != "success" {
		return errors.New("Failed to RemoveDeviceToMultiCastGroup: " + string(response.ResponseBody))
	}
	return nil
}

func (lora *loraService) GetMultiCastGroupDevices(multiCastGroupID string, privateToken string, configs ...CervelloConfigurations) ([]Device, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}
	//http://api.demo.cervello.io/device/v1/lora/organizations/3616505e-0a6d-4893-98f9-18c176775e15/devices?multicastGroupID=00ba2273-0dc7-4d87-94df-46a7bab7bf99

	resource := "/device/v1/lora/organizations/" + config.OrganizationID + "/devices"

	var responseBodyObject getMultiCastGroupDevicesResponse
	response, err := makeHTTPRequest("GET", htppRequest{
		Resource:    resource,
		MetaData:    "GetMultiCastGroupDevices",
		Token:       privateToken,
		QueryParams: QueryParams{Custom: map[string]interface{}{"multicastGroupID": multiCastGroupID}},
	})
	if err != nil {
		internalLog("error", err)
		return []Device{}, err
	}
	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetMultiCastGroupDevices unmarshal response", err)
		return []Device{}, err
	}

	if responseBodyObject.MessageKeys != "success" {
		return []Device{}, errors.New("Failed to GetMultiCastGroupDevices: " + string(response.ResponseBody))
	}

	return responseBodyObject.Result.Devices, nil
}

func (lora *loraService) CreateGateway(gateway LoraGateway, privateToken string, configs ...CervelloConfigurations) (*map[string]interface{}, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}
	// POST /device/v1/lora/organizations/6b099687-0ae0-4df0-9273-a06bd146eb1f/gateways
	// headers [Authorization]
	//
	resource := fmt.Sprintf(
		"/device/v1/lora/organizations/%s/gateways",
		config.OrganizationID,
	)

	var responseBodyObject postLoraGatewayResponse

	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "CreateLoraGateway",
		Token:    privateToken,
		Body:     gateway,
	})

	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "CreateLoraGateway unmarshal response", err)
		return nil, err
	}
	if responseBodyObject.Code != 2001 {
		return nil, errors.New("Failed to CreateLoraGateway: " + string(response.ResponseBody))
	}
	return &responseBodyObject.Result, nil

}

func (lora *loraService) UpdateGateway(id string, gateway LoraGateway, privateToken string, configs ...CervelloConfigurations) (*map[string]interface{}, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}
	// POST /device/v1/lora/organizations/6b099687-0ae0-4df0-9273-a06bd146eb1f/gateways
	// headers [Authorization]
	//
	resource := fmt.Sprintf(
		"/device/v1/lora/organizations/%s/gateways/%s",
		config.OrganizationID,
		id,
	)

	var responseBodyObject postLoraGatewayResponse

	response, err := makeHTTPRequest("PUT", htppRequest{
		Resource: resource,
		MetaData: "UpdateLoraGateway",
		Token:    privateToken,
		Body:     gateway,
	})

	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "UpdateLoraGateway unmarshal response", err)
		return nil, err
	}
	if responseBodyObject.Code != 2002 {
		return nil, errors.New("Failed to UpdateLoraGateway: " + string(response.ResponseBody))
	}
	return &responseBodyObject.Result, nil

}

func (lora *loraService) DeleteGateway(id string, privateToken string, configs ...CervelloConfigurations) (*map[string]interface{}, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}

	// POST /device/v1/lora/organizations/6b099687-0ae0-4df0-9273-a06bd146eb1f/gateways
	// headers [Authorization]
	//
	resource := fmt.Sprintf(
		"/device/v1/lora/organizations/%s/gateways/%s",
		config.OrganizationID,
		id,
	)

	var responseBodyObject postLoraGatewayResponse

	response, err := makeHTTPRequest("DELETE", htppRequest{
		Resource: resource,
		MetaData: "DeleteLoraGateway",
		Token:    privateToken,
	})

	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "DeleteLoraGateway unmarshal response", err)
		return nil, err
	}
	if responseBodyObject.Code != 2004 {
		return nil, errors.New("Failed to DeleteLoraGateway: " + string(response.ResponseBody))
	}
	return &responseBodyObject.Result, nil

}

func (lora *loraService) GetGatewayByID(ID string, privateToken string, configs ...CervelloConfigurations) (*LoraGateway, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}

	// GET http://api.demo.cervello.io/device/v1/lora/organizations/6b099687-0ae0-4df0-9273-a06bd146eb1f/gateways?limit=10&pageSize=10
	resource := fmt.Sprintf(
		"/device/v1/lora/organizations/%s/gateways",
		config.OrganizationID,
		//ID,
	)

	var responseBodyObject *getLoraGatewayByIdResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource: resource,
		MetaData: "GetLoraGatewayByID",
		Token:    privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetLoraGatewayByID unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "unexpected error", string(response.ResponseBody))
	}

	result := LoraGateway{}
	for _, gateway := range responseBodyObject.Result.Data {
		if gateway.ID == ID {
			result = gateway
			break
		}
	}
	if result.ID == "" {
		return nil, errors.New("device not found")
	}

	return &result, nil

}

func (lora *loraService) GetGateways(privateToken string, configs ...CervelloConfigurations) ([]LoraGateway, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}

	// GET http://api.demo.cervello.io/device/v1/lora/organizations/6b099687-0ae0-4df0-9273-a06bd146eb1f/gateways?limit=10&pageSize=10
	resource := fmt.Sprintf(
		"/device/v1/lora/organizations/%s/gateways",
		config.OrganizationID,
		//ID,
	)

	var responseBodyObject *getLoraGatewayByIdResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource: resource,
		MetaData: "GetLoraGatewayByID",
		Token:    privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetLoraGatewayByID unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "unexpected error", string(response.ResponseBody))
	}

	return responseBodyObject.Result.Data, nil

}
