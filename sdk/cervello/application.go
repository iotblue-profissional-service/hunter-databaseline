package cervello

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var (
	// LoraService ...
	ApplicationService applicationServiceInterface = &applicationService{}
)

type applicationService struct {
}

type ApplicationVariableParams struct {
	Key       string      `json:"key"`
	ValueType string      `json:"valueType"`
	Value     interface{} `json:"value"`
}
type ApplicationVariable struct {
	OrganizationID string    `json:"organizationId"`
	ApplicationID  string    `json:"applicationId"`
	Type           string    `json:"type"`
	Key            string    `json:"key"`
	Value          string    `json:"value"`
	ValueType      string    `json:"valueType"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
type ApplicationConfiguration ApplicationVariable

type applicationServiceInterface interface {
	// variable Type can be on of the following Number , String , Boolean
	// if Boolean the Value can be true or false as a string when posting variables or getting
	// if Number the Value is number when posting variables and string when getting variables
	CreateApplicationVariable(variable ApplicationVariableParams, privateToken string, configs ...CervelloConfigurations) error
	// variable Type can be on of the following Number , String , Boolean
	// if Boolean the Value can be true or false as a string when posting variables or getting
	// if Number the Value is number when posting variables and string when getting variables
	CreateBulkApplicationVariables(variable []ApplicationVariableParams, privateToken string, configs ...CervelloConfigurations) error
	// variable Type can be on of the following Number , String , Boolean
	// if Boolean the Value can be true or false as a string when posting variables or getting
	// if Number the Value is number when posting variables and string when getting variables
	CreateApplicationConfiguration(configuration ApplicationVariableParams, privateToken string, configs ...CervelloConfigurations) error
	// variable Type can be on of the following Number , String , Boolean
	// if Boolean the Value can be true or false as a string when posting variables or getting
	// if Number the Value is number when posting variables and string when getting variables
	CreateBulkApplicationConfiguration(configuration []ApplicationVariableParams, privateToken string, configs ...CervelloConfigurations) error
	GetApplicationVariables(queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) ([]ApplicationVariable, error)
	GetApplicationConfigurations(queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) ([]ApplicationConfiguration, error)
}

func (application *applicationService) CreateApplicationVariable(variable ApplicationVariableParams, privateToken string, configs ...CervelloConfigurations) error {
	// POST https://api.release.cervello.io/application/v1/organizations/e4550d32-48f7-4f9c-bdf9-95df8f788140/applications/5dac9016-e1ef-4946-beb1-7f289b0977ca/variables
	// {"key":"exampleVariable","valueType":"String","value":"exampleVariable"}

	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return err
	}

	resource := fmt.Sprintf(
		"/application/v1/organizations/%s/applications/%s/variables",
		config.OrganizationID,
		config.ApplicationID,
	)

	var responseBodyObject baseResponse

	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "CreateApplicationVariable",
		Token:    privateToken,
		Body:     variable,
	})

	if err != nil {
		internalLog("error", err)
		return err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "CreateApplicationVariable unmarshal response", err)
		return err
	}
	if responseBodyObject.Code != 2001 {
		return errors.New("Failed to CreateApplicationVariable: " + string(response.ResponseBody))
	}
	return nil
}

func (application *applicationService) CreateBulkApplicationVariables(variable []ApplicationVariableParams, privateToken string, configs ...CervelloConfigurations) error {
	// POST https://api.release.cervello.io/application/v1/organizations/e4550d32-48f7-4f9c-bdf9-95df8f788140/applications/5dac9016-e1ef-4946-beb1-7f289b0977ca/variables/bulk
	// [{"key":"exampleVariable","valueType":"String","value":"exampleVariable"}]

	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return err
	}

	resource := fmt.Sprintf(
		"/application/v1/organizations/%s/applications/%s/variables/bulk",
		config.OrganizationID,
		config.ApplicationID,
	)

	newVariableMap := make(map[string]interface{})
	for _, tempVar := range variable {
		newVariableMap[tempVar.Key] = tempVar.Value
	}

	// must get old variables first because bulk removes old ones so old one must be sent again in the bulk request
	oldVariables, err := application.GetApplicationVariables(QueryParams{PaginationObj: Pagination{PageNumber: 1, PageSize: 5000}}, privateToken)
	if err != nil {
		return errors.New("Failed to get Old Variables: " + err.Error())
	}

	// append old variables to new variables
	for _, oldVariable := range oldVariables {
		if _, ok := newVariableMap[oldVariable.Key]; !ok {

			tempVariable := ApplicationVariableParams{
				Key:       oldVariable.Key,
				ValueType: oldVariable.ValueType,
				Value:     oldVariable.Value,
			}
			variable = append(variable, tempVariable)
		}
	}

	var responseBodyObject baseResponse

	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "CreateBulkApplicationVariables",
		Token:    privateToken,
		Body:     variable,
	})

	if err != nil {
		internalLog("error", err)
		return err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "CreateBulkApplicationVariables unmarshal response", err)
		return err
	}
	if responseBodyObject.Code != 2001 {
		return errors.New("Failed to CreateBulkApplicationVariables: " + string(response.ResponseBody))
	}
	return nil
}

func (application *applicationService) CreateApplicationConfiguration(configuration ApplicationVariableParams, privateToken string, configs ...CervelloConfigurations) error {
	// POST https://api.release.cervello.io/application/v1/organizations/e4550d32-48f7-4f9c-bdf9-95df8f788140/applications/5dac9016-e1ef-4946-beb1-7f289b0977ca/configurations
	// {"key":"exampleVariable","valueType":"String","value":"exampleVariable"}
	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return err
	}

	resource := fmt.Sprintf(
		"/application/v1/organizations/%s/applications/%s/configurations",
		config.OrganizationID,
		config.ApplicationID,
	)

	var responseBodyObject baseResponse

	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "CreateApplicationConfiguration",
		Token:    privateToken,
		Body:     configuration,
	})

	if err != nil {
		internalLog("error", err)
		return err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "CreateApplicationConfiguration unmarshal response", err)
		return err
	}
	if responseBodyObject.Code != 2001 {
		return errors.New("Failed to CreateApplicationConfiguration: " + string(response.ResponseBody))
	}
	return nil
}

func (application *applicationService) CreateBulkApplicationConfiguration(configuration []ApplicationVariableParams, privateToken string, configs ...CervelloConfigurations) error {
	// POST https://api.release.cervello.io/application/v1/organizations/e4550d32-48f7-4f9c-bdf9-95df8f788140/applications/5dac9016-e1ef-4946-beb1-7f289b0977ca/configurations/bulk
	// [{"key":"exampleVariable","valueType":"String","value":"exampleVariable"}]

	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return err
	}

	resource := fmt.Sprintf(
		"/application/v1/organizations/%s/applications/%s/configurations/bulk",
		config.OrganizationID,
		config.ApplicationID,
	)

	// new configuration map to copare with old configuration
	newConfigMap := make(map[string]interface{})
	for _, config := range configuration {
		newConfigMap[config.Key] = config.Value
	}

	// must get old oldVConfigurations first because bulk removes old ones so old one must be sent again in the bulk request
	oldConfigurations, err := application.GetApplicationVariables(QueryParams{PaginationObj: Pagination{PageNumber: 1, PageSize: 5000}}, privateToken)
	if err != nil {
		return errors.New("Failed to get Old Variables: " + err.Error())
	}

	// append old variables to new Configurations
	for _, oldConfiguration := range oldConfigurations {
		if _, ok := newConfigMap[oldConfiguration.Key]; !ok {
			//do something here
			tempConfiguration := ApplicationVariableParams{
				Key:       oldConfiguration.Key,
				ValueType: oldConfiguration.ValueType,
				Value:     oldConfiguration.Value,
			}
			configuration = append(configuration, tempConfiguration)
		}

	}
	var responseBodyObject baseResponse

	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "CreateBulkApplicationConfiguration",
		Token:    privateToken,
		Body:     configuration,
	})

	if err != nil {
		internalLog("error", err)
		return err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "CreateBulkApplicationConfiguration unmarshal response", err)
		return err
	}
	if responseBodyObject.Code != 2001 {
		return errors.New("Failed to CreateBulkApplicationConfiguration: " + string(response.ResponseBody))
	}
	return nil
}
func (application *applicationService) GetApplicationVariables(queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) ([]ApplicationVariable, error) {
	// POST https://api.release.cervello.io/application/v1/organizations/e4550d32-48f7-4f9c-bdf9-95df8f788140/applications/5dac9016-e1ef-4946-beb1-7f289b0977ca/variables?pageSize=50000
	// [{"key":"exampleVariable","valueType":"String","value":"exampleVariable"}]

	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return nil, err
	}

	resource := fmt.Sprintf(
		"/application/v1/organizations/%s/applications/%s/variables",
		config.OrganizationID,
		config.ApplicationID,
	)

	var responseBodyObject getApplicationVariablesResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource:    resource,
		MetaData:    "GetApplicationVariables",
		Token:       privateToken,
		QueryParams: queryParam,
	})

	if err != nil {
		internalLog("error", err)
		return []ApplicationVariable{}, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetApplicationVariables unmarshal response", err)
		return []ApplicationVariable{}, err
	}
	if responseBodyObject.Code != 2003 {
		return []ApplicationVariable{}, errors.New("Failed to GetApplicationVariables: " + string(response.ResponseBody))
	}
	return responseBodyObject.Result.ApplicationVariables, nil
}

func (application *applicationService) GetApplicationConfigurations(queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) ([]ApplicationConfiguration, error) {
	// POST https://api.release.cervello.io/application/v1/organizations/e4550d32-48f7-4f9c-bdf9-95df8f788140/applications/5dac9016-e1ef-4946-beb1-7f289b0977ca/configurations?pageSize=50000
	// [{"key":"exampleVariable","valueType":"String","value":"exampleVariable"}]

	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return nil, err
	}

	resource := fmt.Sprintf(
		"/application/v1/organizations/%s/applications/%s/configurations",
		config.OrganizationID,
		config.ApplicationID,
	)

	var responseBodyObject getApplicationConfigurationsResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource:    resource,
		MetaData:    "GetApplicationConfigurations",
		Token:       privateToken,
		QueryParams: queryParam,
	})

	if err != nil {
		internalLog("error", err)
		return []ApplicationConfiguration{}, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetApplicationConfigurations unmarshal response", err)
		return []ApplicationConfiguration{}, err
	}
	if responseBodyObject.Code != 2003 {
		return []ApplicationConfiguration{}, errors.New("Failed to GetApplicationConfigurations: " + string(response.ResponseBody))
	}
	return responseBodyObject.Result.ApplicationConfiguration, nil
}
