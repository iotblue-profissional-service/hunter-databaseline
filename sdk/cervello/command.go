package cervello

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var (
	// CommandService ...
	CommandService commandServiceInterface = &commandService{}
)

type commandService struct {
}

type commandServiceInterface interface {
	// get commands under a device
	GetDeviceCommands(deviceID string, queryParams QueryParams, privateToken string, configs ...CervelloConfigurations) ([]Command, error)
	// GetCommandTemplates ...
	GetCommandTemplates(queryParams QueryParams, privateToken string, configs ...CervelloConfigurations) ([]CommandTemplate, error)
	// CreateDeviceCommand ...
	CreateDeviceCommand(deviceID string, command Command, privateToken string, configs ...CervelloConfigurations) (*Command, error)
	DeleteDeviceCommand(deviceID string, commandID string, privateToken string, configs ...CervelloConfigurations) error
	ExecuteDeviceCommand(params ExecuteDeviceCommandParamters, privateToken string, configs ...CervelloConfigurations) error
	CreateCommandTemplate(commandTemplate CommandTemplate, privateToken string, configs ...CervelloConfigurations) (*CommandTemplate, error)
	UpdateCommandTemplate()
	ExecuteCommandTemplate(params ExecuteCommandTemplateParamters, privateToken string, configs ...CervelloConfigurations) error
	DeleteCommandTemplate(commandTemplateID string, privateToken string, configs ...CervelloConfigurations) error
}

// Command data type
type Command struct {
	ID                    string      `json:"id,omitempty"`
	OrganizationID        string      `json:"organizationId,omitempty"`
	DeviceID              string      `json:"deviceId,omitempty"`
	Name                  string      `json:"name,omitempty"`
	IsActive              bool        `json:"isActive,omitempty"`
	Command               string      `json:"command,omitempty"`
	Parameters            interface{} `json:"parameters,omitempty"`
	CommunicationProtocol interface{} `json:"communicationProtocol,omitempty"`
	Type                  string      `json:"type,omitempty"`
	CreatedAt             time.Time   `json:"createdAt,omitempty"`
	UpdatedAt             time.Time   `json:"updatedAt,omitempty"`
	Protocol              string      `json:"protocol,omitempty"`
	ProtocolConfiguration struct {
		ProtocolType string `json:"protocolType,omitempty"`
	} `json:"protocolConfiguration,omitempty"`
}

// CommandTemplate ...
type CommandTemplate struct {
	ID                    string    `json:"id,omitempty"`
	OrganizationID        string    `json:"organizationId,omitempty"`
	Name                  string    `json:"name,omitempty"`
	Commands              []Command `json:"commands,omitempty"`
	Tags                  []string  `json:"tags,omitempty"`
	IsActive              bool      `json:"isActive,omitempty"`
	CommunicationProtocol struct {
		Name           string `json:"name,omitempty"`
		Configurations struct {
			ProtocolType string `json:"protocolType,omitempty"`
		} `json:"configurations,omitempty"`
	} `json:"communicationProtocol,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Devices   []struct {
		ID string `json:"id,omitempty"`
	} `json:"devices,omitempty"`
}

// ExecuteDeviceCommandParamters .....
type ExecuteDeviceCommandParamters struct {
	DeviceID   string
	CommandID  string
	Timeout    int                    `json:"timeout"`
	Retries    int                    `json:"retries"`
	ClientID   string                 `json:"clientId,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// ExecuteCommandTemplateParamters ...
type ExecuteCommandTemplateParamters struct {
	CommandTemplateID string
	Timeout           int                    `json:"timeout"`
	Retries           int                    `json:"retries"`
	Parameters        map[string]interface{} `json:"parameters"`
	DevicesIds        []string               `json:"devicesIds"`
}

type internalExecuteDeviceCommandParamters struct {
	Timeout    int                    `json:"timeout"`
	Retries    int                    `json:"retries"`
	ClientID   string                 `json:"clientId,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// ModbusCommandParameters ...
type ModbusCommandParameters struct {
	Address  int `json:"address,omitempty"`
	Sequence int `json:"sequence,omitempty"`
	Value    int `json:"value,omitempty"`
}

// ModbusCommandCommunicationProtocol ...
type ModbusCommandCommunicationProtocol struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Name     string `json:"name,omitempty"`
	SlaveID  int    `json:"slaveId,omitempty"`
	Address  int    `json:"address,omitempty"`
	Sequence int    `json:"sequence,omitempty"`
}

//MQTTCommandCommunicationProtocol .....
type MQTTCommandCommunicationProtocol struct {
	Port           int                       `json:"port,omitempty"`
	Name           string                    `json:"name,omitempty"`
	Configurations MQQTCommandConfigurations `json:"configurations,omitempty"`
}

// MQQTCommandConfigurations ...
type MQQTCommandConfigurations struct {
	Retain bool `json:"retain,omitempty"`
	Qos    int  `json:"qos,omitempty"`
}

// GetDeviceCommands ...
func (s *commandService) GetDeviceCommands(deviceID string, queryParams QueryParams, privateToken string, configs ...CervelloConfigurations) ([]Command, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}
	//https://api.staging.cervello.io/device/v1/organizations/8bd3d7dc-1b71-4ae7-9d23-213f7e83a102/devices/62e56ab9-573c-4cf0-9245-d48455f2a436/commands
	// GetDeviceByID ...
	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/devices/%s/commands",
		config.OrganizationID,
		deviceID,
	)

	var responseBodyObject *getDeviceCommandsResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource:    resource,
		MetaData:    "GetDeviceCommands",
		QueryParams: queryParams,
		Token:       privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetDeviceCommands unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "GetDeviceCommands error", string(response.ResponseBody))
	}

	return responseBodyObject.Result.Commands, nil

}

// GetCommandTemplates ...
func (s *commandService) GetCommandTemplates(queryParams QueryParams, privateToken string, configs ...CervelloConfigurations) ([]CommandTemplate, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}
	//	https://api.staging.cervello.io/device/v1/organizations/8bd3d7dc-1b71-4ae7-9d23-213f7e83a102/commandtemplates/?pageSize=10&offset=0
	// GetDeviceByID ...
	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/commandtemplates",
		config.OrganizationID,
	)

	var responseBodyObject *getCommandTemplatesResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource:    resource,
		MetaData:    "GetCommandTemplates",
		QueryParams: queryParams,
		Token:       privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetCommandTemplates unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "GetCommandTemplates error", string(response.ResponseBody))
	}

	return responseBodyObject.Result.CommandTemplate, nil

}

// CreateDeviceCommand ...
func (s *commandService) CreateDeviceCommand(deviceID string, command Command, privateToken string, configs ...CervelloConfigurations) (*Command, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}
	//https://api.staging.cervello.io/device/v1/organizations/8bd3d7dc-1b71-4ae7-9d23-213f7e83a102/devices/2be08ffe-e4c1-48ac-a670-6fdfaba4e786/commands
	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/devices/%s/commands",
		config.OrganizationID,
		deviceID,
	)

	var responseBodyObject createDeviceCommandResponse

	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "CreateDeviceCommand",
		Token:    privateToken,
		Body:     command,
	})

	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "CreateDeviceCommand unmarshal response", err)
		return nil, err
	}
	if responseBodyObject.Code != 2001 {
		return nil, errors.New("Failed to CreateDeviceCommand: " + string(response.ResponseBody))
	}

	return &responseBodyObject.Result, nil
}

// DeleteDeviceCommand ...
func (s *commandService) DeleteDeviceCommand(deviceID string, commandID string, privateToken string, configs ...CervelloConfigurations) error {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return err
	}
	//https://api.staging.cervello.io/device/v1/organizations/8bd3d7dc-1b71-4ae7-9d23-213f7e83a102/devices/2be08ffe-e4c1-48ac-a670-6fdfaba4e786/commands/e9dbdf32-77c5-4a0b-a4cc-0930ddd0b6d3
	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/devices/%s/commands/%s",
		config.OrganizationID,
		deviceID,
		commandID,
	)

	var responseBodyObject baseResponse

	response, err := makeHTTPRequest("DELETE", htppRequest{
		Resource: resource,
		MetaData: "DeleteDeviceCommand",
		Token:    privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "DeleteDeviceCommand unmarshal response", err)
		return err
	}

	if responseBodyObject.Code != 2004 {
		return errors.New("Failed to DeleteDeviceCommand: " + string(response.ResponseBody))
	}

	return nil
}

// ExecuteDeviceCommand ...
func (s *commandService) ExecuteDeviceCommand(params ExecuteDeviceCommandParamters, privateToken string, configs ...CervelloConfigurations) error {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return err
	}
	//https://api.staging.cervello.io/device/v1/organizations/8bd3d7dc-1b71-4ae7-9d23-213f7e83a102/devices/2be08ffe-e4c1-48ac-a670-6fdfaba4e786/commands/c8198f10-2562-4254-a6d9-1100c480eb98/exec
	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/devices/%s/commands/%s/exec",
		config.OrganizationID,
		params.DeviceID,
		params.CommandID,
	)

	var responseBodyObject baseResponse

	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "ExecuteDeviceCommand",
		Token:    privateToken,
		Body: internalExecuteDeviceCommandParamters{
			Timeout:    params.Timeout,
			Retries:    params.Retries,
			ClientID:   params.ClientID,
			Parameters: params.Parameters,
		},
	})
	if err != nil {
		internalLog("error", err)
		return err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "ExecuteDeviceCommand unmarshal response", err)
		return err
	}

	if responseBodyObject.Code != 2005 {
		return errors.New("Failed to ExecuteDeviceCommand: " + string(response.ResponseBody))
	}

	return nil
}

// CreateCommandTemplate ...
func (s *commandService) CreateCommandTemplate(commandTemplate CommandTemplate, privateToken string, configs ...CervelloConfigurations) (*CommandTemplate, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}
	//https://api.staging.cervello.io/device/v1/organizations/8bd3d7dc-1b71-4ae7-9d23-213f7e83a102/commandtemplates
	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/commandtemplates",
		config.OrganizationID,
	)

	var responseBodyObject createCommandTemplatesResponse

	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "CreateCommandTemplate",
		Token:    privateToken,
		Body:     commandTemplate,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "CreateCommandTemplate unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2001 {
		return nil, errors.New("Failed to CreateCommandTemplate: " + string(response.ResponseBody))
	}

	return &responseBodyObject.Result, nil
}

// UpdateCommandTemplate ...
func (s *commandService) UpdateCommandTemplate() {

}

// DeleteCommandTemplate ...
func (s *commandService) DeleteCommandTemplate(commandTemplateID string, privateToken string, configs ...CervelloConfigurations) error {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return err
	}
	//https://api.staging.cervello.io/device/v1/organizations/8bd3d7dc-1b71-4ae7-9d23-213f7e83a102/commandtemplates/2c2a4e51-1328-40d0-b4ee-0e3cdf7d7c0c
	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/commandtemplates/%s",
		config.OrganizationID,
		commandTemplateID,
	)

	var responseBodyObject baseResponse

	response, err := makeHTTPRequest("DELETE", htppRequest{
		Resource: resource,
		MetaData: "DeleteCommandTemplate",
		Token:    privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "DeleteCommandTemplate unmarshal response", err)
		return err
	}

	if responseBodyObject.Code != 2004 {
		return errors.New("Failed to DeleteCommandTemplate: " + string(response.ResponseBody))
	}

	return nil
}

// ExecuteCommandTemplate ...
func (s *commandService) ExecuteCommandTemplate(params ExecuteCommandTemplateParamters, privateToken string, configs ...CervelloConfigurations) error {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return err
	}
	//https: //api.staging.cervello.io/device/v1/organizations/8bd3d7dc-1b71-4ae7-9d23-213f7e83a102/commandtemplates/fc66b5a6-9197-4706-92a8-f6e0a2411344/exec
	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/commandtemplates/%s/exec",
		config.OrganizationID,
		params.CommandTemplateID,
	)

	var responseBodyObject baseResponse

	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "ExecuteCommandTemplate",
		Token:    privateToken,
		Body:     params,
	})
	if err != nil {
		internalLog("error", err)
		return err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "ExecuteCommandTemplate unmarshal response", err)
		return err
	}
	if responseBodyObject.Code != 2005 {
		return errors.New("Failed to ExecuteCommandTemplate: " + string(response.ResponseBody))
	}

	return nil
}
