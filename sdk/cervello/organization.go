package cervello

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var (
	// CommandService ...
	OrganizationService organizationServiceInterface = &organizationService{}
)

type organizationService struct {
}

type organizationServiceInterface interface {
	CreateCredentail(privateToken string, configs ...CervelloConfigurations) (*OrganizationCredential, error)
	GetCredentail(queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) ([]OrganizationCredential, error)
	DeleteCredentail(credentialId, privateToken string, configs ...CervelloConfigurations) error
	GetOrganisations(queryParam QueryParams, privateToken string) ([]Organization, error)
	GetOrganizationById(organizationId string, privateToken string) (*Organization, error)
}

type OrganizationCredential struct {
	AccessKey      string `json:"accessKey,omitempty"`
	AccessToken    string `json:"accessToken,omitempty"`
	OrganizationID string `json:"organizationId,omitempty"`
	Id             string `json:"id,omitempty"`
	CreatedAt      string `json:"createdAt,omitempty"`
	UpdatedAt      string `json:"updatedAt,omitempty"`
}

type Organization struct {
	ID                  string    `json:"id"`
	OwnerID             string    `json:"ownerId"`
	Description         string    `json:"description"`
	Status              string    `json:"status"`
	Title               string    `json:"title"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
	OrganizationMembers []struct {
		Role string `json:"role"`
	} `json:"organizationMembers"`
}

// CreateCredentail ... used to create organisation credentail
func (s *organizationService) GetOrganisations(queryParam QueryParams, privateToken string) ([]Organization, error) {
	//GET
	//https: http://api.demo.cervello.io/organization/v1/organizations?pageSize=10
	resource := fmt.Sprintf(
		"/organization/v1/organizations",
	)

	var responseBodyObject getOrganisationsResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource: resource,
		MetaData: "GetOrganisations",
		Token:    privateToken,
	})

	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	// log.Printf(string(response.ResponseBody))
	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetOrganisations unmarshal response", err)
		return nil, err
	}
	if responseBodyObject.Code != 2003 {
		return nil, errors.New("Failed to GetOrganisations: " + string(response.ResponseBody))
	}

	return responseBodyObject.Result.Data, nil
}

// GetOrganizationById is used to get organization by ID
func (s *organizationService) GetOrganizationById(organizationId string, privateToken string) (*Organization, error) {
	//GET
	//https: http://api.demo.cervello.io/organization/v1/organizations/8bd3d7dc-1b71-4ae7-9d23-213f7e83a102
	resource := fmt.Sprintf("/organization/v1/organizations/%s", organizationId)
	var responseBodyObject getOrganisationResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource: resource,
		MetaData: "GetOrganizationById",
		Token:    privateToken,
	})

	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	// log.Printf(string(response.ResponseBody))
	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetOrganisation unmarshal response", err)
		return nil, err
	}
	if responseBodyObject.Code != 2003 {
		return nil, errors.New("Failed to GetOrganisation: " + string(response.ResponseBody))
	}

	return &responseBodyObject.Result.Data, nil
}

// CreateCredentail ... used to create organisation credentail
func (s *organizationService) CreateCredentail(privateToken string, configs ...CervelloConfigurations) (*OrganizationCredential, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}
	//GET
	//https: //api.staging.cervello.io/device/v1/organizations/8bd3d7dc-1b71-4ae7-9d23-213f7e83a102/credentials	// GetDeviceByID ...
	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/credentials",
		config.OrganizationID,
	)

	var responseBodyObject createOrganizationCredentialResponse

	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "CreateOrganisationCredentail",
		Token:    privateToken,
	})

	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	// log.Printf(string(response.ResponseBody))
	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "CreateOrganisationCredentail unmarshal response", err)
		return nil, err
	}
	if responseBodyObject.Code != 2001 {
		return nil, errors.New("Failed to CreateOrganisationCredentail: " + string(response.ResponseBody))
	}

	return &responseBodyObject.Result, nil
}

// GetCredentail ... used to create organisation credentail
func (s *organizationService) GetCredentail(queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) ([]OrganizationCredential, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}
	//GET
	//https://api.staging.cervello.io/device/v1/organizations/8bd3d7dc-1b71-4ae7-9d23-213f7e83a102/credentials	resource := fmt.Sprintf(
	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/credentials",
		config.OrganizationID,
	)

	var responseBodyObject getOrganizationCredentialResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource:    resource,
		MetaData:    "GetOrganisationCredentail",
		Token:       privateToken,
		QueryParams: queryParam,
	})

	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetOrganisationCredentail unmarshal response", err)
		return nil, err
	}
	if responseBodyObject.Code != 2003 {
		return nil, errors.New("Failed to GetOrganisationCredentail: " + string(response.ResponseBody))
	}

	return responseBodyObject.Result.Credentails, nil

}

// DeleteCredentail ... used to create delete organisation credentail
func (s *organizationService) DeleteCredentail(credentialId, privateToken string, configs ...CervelloConfigurations) error {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return err
	}
	//https: //api.staging.cervello.io/device/v1/organizations/8bd3d7dc-1b71-4ae7-9d23-213f7e83a102/credentials/9451bcd9-188e-440c-8227-0ae31966d620
	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/credentials/%s",
		config.OrganizationID,
		credentialId,
	)

	var responseBodyObject baseResponse

	response, err := makeHTTPRequest("DELETE", htppRequest{
		Resource: resource,
		MetaData: "deleteOrganisationCredentail",
		Token:    privateToken,
	})

	if err != nil {
		internalLog("error", err)
		return err
	}
	// log.Printf(string(response.ResponseBody))
	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "deleteOrganisationCredentail unmarshal response", err)
		return err
	}
	if responseBodyObject.Code != 2005 {
		return errors.New("Failed to deleteOrganisationCredentail: " + string(response.ResponseBody))
	}

	return nil

}
