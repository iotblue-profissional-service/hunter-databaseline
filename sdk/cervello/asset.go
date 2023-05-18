package cervello

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Asset struct {
	ID             string `json:"id,omitempty"`
	AssetType      string `json:"assetType,omitempty"`
	Name           string `json:"name,omitempty"`
	ReferenceName  string `json:"referenceName,omitempty"`
	AdditionalInfo string `json:"additionalInfo,omitempty"`
	ParentAsset    []struct {
		AssetId string `json:"assetId"`
	} `json:"parentAsset"`
	CustomFields map[string]interface{} `json:"customFields,omitempty"`
	CreatedAt    string                 `json:"createdAt,omitempty"`
	UpdatedAt    string                 `json:"updatedAt,omitempty"`
}

type ChildAsset struct {
	AssetId        string                 `json:"assetId,omitempty"`
	ResourceId     string                 `json:"resourceId,omitempty"`
	AssetType      string                 `json:"assetType,omitempty"`
	Name           string                 `json:"name,omitempty"`
	ReferenceName  string                 `json:"referenceName,omitempty"`
	AdditionalInfo string                 `json:"additionalInfo,omitempty"`
	CustomFields   map[string]interface{} `json:"customFields,omitempty"`
}

type AssetDevice struct {
	ID         string   `json:"id,omitempty"`
	DeviceType string   `json:"deviceType,omitempty"`
	Name       string   `json:"name,omitempty"`
	Tags       []string `json:"tags,omitempty"`
}

type AssetParentAsset struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Depth int    `json:"depth,omitempty"`
}

func GetAsset(referenceID string, privateToken string, configs ...CervelloConfigurations) (*Asset, error) {
	// GET https://api.staging.cervello.io/application/v1/organizations/0feb9329-cf80-4c61-812b-c67eb74d5c33/applications/de1be3f0-6b43-47e8-8c23-984b07b0bf51/assets/
	// ?pageSize=1&filters[0][key]=referenceName&filters[0][operator]=eq&filters[0][value]=2388sa
	// headers [Authorization: "Bearer {token}"]
	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return nil, err
	}

	resource := fmt.Sprintf(
		"/application/v1/organizations/%s/applications/%s/assets",
		config.OrganizationID,
		config.ApplicationID,
	)

	var responseBodyObject *getAssetResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource: resource,
		MetaData: "GetAsset",
		QueryParams: QueryParams{
			Filters: []Filter{
				{Key: "referenceName", Op: "eq", Value: referenceID},
			},
			PaginationObj: Pagination{PageNumber: 1, PageSize: 1},
		},
		Token: privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetAsset unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "GetAsset error", string(response.ResponseBody))
	}
	if responseBodyObject.Result.Assets == nil || len(responseBodyObject.Result.Assets) == 0 {
		return &Asset{}, nil
	}
	return &responseBodyObject.Result.Assets[0], nil

}

func GetAssetByID(assetID string, privateToken string, configs ...CervelloConfigurations) (*Asset, error) {
	//https://api.release.cervello.io/application/v1/organizations/5e06127e-0aed-4a41-892f-2c22f59ceece/applications/243daa4c-ecb6-4db2-861d-7551af97b9cf/assets/3c6ab502-30fd-43ff-baef-e2455eecbd5f
	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return nil, err
	}

	resource := fmt.Sprintf(
		"/application/v1/organizations/%s/applications/%s/assets/%s",
		config.OrganizationID,
		config.ApplicationID,
		assetID,
	)

	var responseBodyObject *getAssetByIDResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource: resource,
		MetaData: "GetAssetByID",
		Token:    privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetAssetByID unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "GetAssetByID error", string(response.ResponseBody))
	}

	return &responseBodyObject.Result, nil

}

func GetAssetAttributes(assetID string, privateToken string, configs ...CervelloConfigurations) (map[string]interface{}, error) {
	// GET https://api.staging.cervello.io/data/v1/organizations/9a450c80-8511-4de5-8a1c-a3ecc6e9c862/applications/535880fd-91a5-4c2f-b7f6-4fe21869c730/assets/c428a94e-30bd-44e3-9970-b4c1ee558321/attributes
	// ?pageSize=100
	// headers [Authorization: "Bearer {token}"]

	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return nil, err
	}

	resource := fmt.Sprintf(
		"/data/v1/organizations/%s/applications/%s/assets/%s/attributes",
		config.OrganizationID,
		config.ApplicationID,
		assetID,
	)

	responseBodyObject := new(map[string]interface{})

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource: resource,
		MetaData: "GetAssetAttributes",
		Token:    privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetAssetAttributes unmarshal response", err)
		return nil, err
	}

	return *responseBodyObject, nil
}

func CreateAsset(a Asset, privateToken string, configs ...CervelloConfigurations) (*Asset, error) {
	// POST https://api.staging.cervello.io/
	// application/v1/
	// organizations/dbb02a67-6ab6-4485-b955-e485ade0a003/
	// applications/7d12790c-5bbd-4168-9689-34436ebd8543/
	// assets
	// {"name":"building01","assetType":"building", referenceName: ""}
	// headers [Authorization]

	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return nil, err
	}

	resource := fmt.Sprintf(
		"/application/v1/organizations/%s/applications/%s/assets",
		config.OrganizationID,
		config.ApplicationID,
	)

	var responseBodyObject postAssetResponse

	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "CreateAsset",
		Token:    privateToken,
		Body:     a,
	})

	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "CreateAsset unmarshal response", err)
		return nil, err
	}
	if responseBodyObject.Code != 2001 {
		return nil, errors.New("Failed to CreateAsset: " + string(response.ResponseBody))
	}

	return &responseBodyObject.Result, nil
}

func UpdateAsset(assetID string, asset Asset, privateToken string, configs ...CervelloConfigurations) error {
	// PUT https://api.staging.cervello.io/application/v1/organizations/9a450c80-8511-4de5-8a1c-a3ecc6e9c862/applications/535880fd-91a5-4c2f-b7f6-4fe21869c730/assets/d8dd34c9-9bd8-48eb-b99f-9e15ce50ee49
	// headers [Authorization]
	//
	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return err
	}

	resource := fmt.Sprintf(
		"/application/v1/organizations/%s/applications/%s/assets/%s",
		config.OrganizationID,
		config.ApplicationID,
		assetID,
	)

	var responseBodyObject baseResponse
	response, err := makeHTTPRequest("PUT", htppRequest{
		Resource: resource,
		MetaData: "UpdateAsset",
		Body:     asset,
		Token:    privateToken,
	})

	if err != nil {
		internalLog("error", err)
		return err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "UpdateAsset unmarshal response", err)
		return err
	}

	if responseBodyObject.Code != 2002 {
		return errors.New("Failed to update Asset: " + string(response.ResponseBody))
	}
	return nil
}

func CreateAssetAttributes(assetID string, attributes map[string]interface{}, privateToken string, configs ...CervelloConfigurations) (bool, error) {
	// POST https://api.staging.cervello.io/
	// data/v1/
	// organizations/dbb02a67-6ab6-4485-b955-e485ade0a003/
	// applications/7d12790c-5bbd-4168-9689-34436ebd8543/
	// assets/4216422e-36c3-4987-81f5-739b5a77b2d7/
	// attributes
	// {"data":{"attrib01":"0111"}}
	// headers [Authorization]

	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return false, err
	}

	resource := fmt.Sprintf(
		"/data/v1/organizations/%s/applications/%s/assets/%s/attributes",
		config.OrganizationID,
		config.ApplicationID,
		assetID,
	)

	var attributeBody = map[string]interface{}{
		"data": attributes,
	}

	var responseBodyObject baseResponse

	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "CreateAssetAttributes",
		Token:    privateToken,
		Body:     attributeBody,
	})

	if err != nil {
		internalLog("error", err)
		return false, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "CreateAssetAttributes unmarshal response", err)
		return false, err
	}
	if responseBodyObject.Code != 2001 {
		return false, errors.New("Failed to CreateAssetAttributes: " + string(response.ResponseBody))
	}

	return true, nil

}

func AssignDeviceToAsset(deviceID string, assetID string, privateToken string, configs ...CervelloConfigurations) (bool, error) {
	// POST https://api.staging.cervello.io/
	// 	compose/v1/
	// 	organizations/dbb02a67-6ab6-4485-b955-e485ade0a003/
	// 	applications/7d12790c-5bbd-4168-9689-34436ebd8543/
	// 	devices/055685c4-06dc-4996-b36c-7989c354ef66/
	// 	assets
	// ["4216422e-36c3-4987-81f5-739b5a77b2d7"] //asset IDs
	// headers [Authorization]

	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return false, err
	}

	resource := fmt.Sprintf(
		"/compose/v1/organizations/%s/applications/%s/devices/%s/assets",
		config.OrganizationID,
		config.ApplicationID,
		deviceID,
	)

	////////////
	var assetsIDs = []string{
		assetID,
	}
	var responseBodyObject baseResponse

	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "AssignDeviceToAsset",
		Token:    privateToken,
		Body:     assetsIDs,
	})

	if err != nil {
		internalLog("error", err)
		return false, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "AssignDeviceToAsset unmarshal response", err)
		return false, err
	}
	if responseBodyObject.Code != 2001 {
		return false, errors.New("Failed to AssignDeviceToAsset: " + string(response.ResponseBody))
	}

	return true, nil

}

func AssignAssetToAsset(assetID string, parentAssetID string, privateToken string, configs ...CervelloConfigurations) (bool, error) {
	// POST https://api.staging.cervello.io/
	// application/v1/
	// organizations/dbb02a67-6ab6-4485-b955-e485ade0a003/
	// applications/7d12790c-5bbd-4168-9689-34436ebd8543/
	// assets/da61eea5-9553-4af9-bb7c-3eaf70e9876d/
	// parents
	// ["4216422e-36c3-4987-81f5-739b5a77b2d7"] //parent asset IDs
	// headers [Authorization]

	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return false, err
	}

	resource := fmt.Sprintf(
		"/application/v1/organizations/%s/applications/%s/assets/%s/parents",
		config.OrganizationID,
		config.ApplicationID,
		assetID,
	)

	var parentAssetsIDs = []string{
		parentAssetID,
	}
	var responseBodyObject baseResponse

	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "AssignAssetToAsset",
		Token:    privateToken,
		Body:     parentAssetsIDs,
	})

	if err != nil {
		internalLog("error", err)
		return false, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "AssignAssetToAsset unmarshal response", err)
		return false, err
	}
	if responseBodyObject.Code != 2001 {
		return false, errors.New("Failed to AssignAssetToAsset: " + string(response.ResponseBody))
	}

	return true, nil

}

// GetAssetDevicesFiltered ... status = connected | disconnected | all
//
// to get connected devices use filter lastConnectionStatus eq true
//
// to get all devices send empty filters
func GetAssetDevicesFiltered(assetID string, queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) ([]Device, error) {

	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return nil, err
	}

	resource := fmt.Sprintf(
		"/compose/v1/organizations/%s/applications/%s/assets/%s/devices",
		config.OrganizationID,
		config.ApplicationID,
		assetID,
	)

	var responseBodyObject getAssetDevicesResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource:    resource,
		MetaData:    "GetAssetDevicesFiltered",
		QueryParams: queryParam,
		Token:       privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetAssetDevicesFiltered unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "GetAssetDevicesFiltered error", string(response.ResponseBody))
	}

	return responseBodyObject.Result.AssetDevices, nil

}

// GetAssetDevices ... status = connected | disconnected | all
func GetAssetDevices(assetID string, status string, queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) ([]Device, error) {
	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return nil, err
	}

	resource := fmt.Sprintf(
		"/compose/v1/organizations/%s/applications/%s/assets/%s/devices",
		config.OrganizationID,
		config.ApplicationID,
		assetID,
	)

	data := url.Values{}

	data = addQueryParamsToRequest(data, queryParam)

	switch status {
	case "connected":
		data.Set("filters[0][key]", "lastConnectionStatus")
		data.Set("filters[0][operator]", "eq")
		data.Set("filters[0][value]", "true")
	case "disconnected":
		data.Set("filters[0][key]", "lastConnectionStatus")
		data.Set("filters[0][operator]", "eq")
		data.Set("filters[0][value]", "false")
	default:
	}

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
		internalLog("error", "GetAssetDevices-NewRequest: ")
		internalLog("error", err)
		return nil, err
	}

	addAuthHeader(req, privateToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		internalLog("error", "GetAssetDevices-Do: ")
		internalLog("error", err)
		return nil, err
	}

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		internalLog("error", "GetAssetDevices-ReadAll: ")
		internalLog("error", err)
		return nil, err
	}

	resp.Body.Close()
	// Assets := new(getAssetResponse)
	var AssetDevices getAssetDevicesResponse

	err = json.Unmarshal(f, &AssetDevices)
	if err != nil {
		internalLog("error", "GetAssetDevices-Unmarshal: ")
		internalLog("error", err)
		return nil, err
	}

	if AssetDevices.Code != 2003 {
		return nil, err
	}

	return AssetDevices.Result.AssetDevices, nil
}

//GetAssetsFlitered ...
//
// to get Assets by assetType set filter assetType eq pole
func GetAssetsFlitered(queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) ([]Asset, error) {
	//https://api.release.cervello.io/application/v1/organizations/724a3eb5-6a76-4fb6-9a6d-499491e9c1d7/applications/b3c4692e-e715-4098-a683-a9a238dd3d00/assets/?filters[0][key]=assetType&filters[0][operator]=eq&filters[0][value]=zone

	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return nil, err
	}

	resource := fmt.Sprintf(
		"/application/v1/organizations/%s/applications/%s/assets",
		config.OrganizationID,
		config.ApplicationID,
	)

	var responseBodyObject getAssetResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource:    resource,
		MetaData:    "GetAssetsFlitered",
		QueryParams: queryParam,
		Token:       privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetAssetsFlitered unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "GetAssetsFlitered error", string(response.ResponseBody))
	}

	return responseBodyObject.Result.Assets, nil

}

func GetAssetsByAssetType(assetType string, queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) ([]Asset, error) {
	//https://api.release.cervello.io/application/v1/organizations/724a3eb5-6a76-4fb6-9a6d-499491e9c1d7/applications/b3c4692e-e715-4098-a683-a9a238dd3d00/assets/?filters[0][key]=assetType&filters[0][operator]=eq&filters[0][value]=zone

	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return nil, err
	}

	resource := fmt.Sprintf(
		"/application/v1/organizations/%s/applications/%s/assets",
		config.OrganizationID,
		config.ApplicationID,
	)

	data := url.Values{}
	data.Set("filters[0][key]", "assetType")
	data.Set("filters[0][operator]", "eq")
	data.Set("filters[0][value]", assetType)
	data = addQueryParamsToRequest(data, queryParam)

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
		internalLog("error", "GetAssetsByAssetType-NewRequest: ")
		internalLog("error", err)
		return nil, err
	}

	addAuthHeader(req, privateToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		internalLog("error", "GetAssetsByAssetType-Do: ")
		internalLog("error", err)
		return nil, err
	}

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		internalLog("error", "GetAssetsByAssetType-ReadAll: ")
		internalLog("error", err)
		return nil, err
	}

	resp.Body.Close()
	// Assets := new(getAssetResponse)
	var Assets getAssetResponse

	err = json.Unmarshal(f, &Assets)
	if err != nil {
		internalLog("error", "GetAssetsByAssetType-Unmarshal: ")
		internalLog("error", err, string(f))
		return nil, err
	}

	if Assets.Code != 2003 {
		return nil, err
	}

	return Assets.Result.Assets, nil
}

func CountAssetsByAssetType(assetType string, privateToken string, configs ...CervelloConfigurations) (int, error) {

	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return 0, err
	}

	resource := fmt.Sprintf(
		"/application/v1/organizations/%s/applications/%s/assets",
		config.OrganizationID,
		config.ApplicationID,
	)

	var responseBodyObject getAssetResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource: resource,
		MetaData: "CountAssetsByAssetType",
		QueryParams: QueryParams{
			Filters: []Filter{
				{Key: "assetType", Op: "eq", Value: assetType},
			},
			PaginationObj: Pagination{PageNumber: 1, PageSize: 999999999999999999},
		},
		Token: privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return 0, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "CountAssetsByAssetType unmarshal response", err)
		return 0, err
	}

	if responseBodyObject.Code != 2003 {
		return 0, fmt.Errorf("%v-%v", "CountAssetsByAssetType error", string(response.ResponseBody))
	}

	return int(responseBodyObject.Result.TotalCount.(float64)), nil

}

//GetAssetLastTelmetries .... key must be empty for all keys
func GetAssetLastTelmetries(assetID string, key string, queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) ([]Telmetry, error) {
	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return nil, err
	}

	//https://api.release.cervello.io/compose/v1/organizations/6df9aab1-0a16-4c24-9176-6a879a5a514e/applications/c033e5c8-50d7-42b3-9e91-153ec2416510/assets/2c81e57a-a357-44e9-9ed9-c30b6de0317a/telemetries?aggregation=last&filters[0][key]=key&filters[0][operator]=eq&filters[0][value]=key
	resource := fmt.Sprintf(
		"/compose/v1/organizations/%s/applications/%s/assets/%s/telemetries",
		config.OrganizationID,
		config.ApplicationID,
		assetID,
	)

	data := url.Values{}
	data.Set("aggregation", "last")
	data = addQueryParamsToRequest(data, queryParam)

	switch key {
	case "":
	default:
		data.Set("filters[0][key]", "key")
		data.Set("filters[0][operator]", "eq")
		data.Set("filters[0][value]", key)
	}

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
		internalLog("error", "GetAssetLastTelmetries-NewRequest: ")
		internalLog("error", err)
		return []Telmetry{}, err
	}

	addAuthHeader(req, privateToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		internalLog("error", "GetAssetLastTelmetries-Do: ")
		internalLog("error", err)
		return []Telmetry{}, err
	}

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		internalLog("error", "GetAssetLastTelmetries-ReadAll: ")
		internalLog("error", err)
		return []Telmetry{}, err
	}

	resp.Body.Close()

	// Assets := new(getAssetResponse)
	var tagTelemetries getTagTelemtriesResponse

	err = json.Unmarshal(f, &tagTelemetries)
	if err != nil {
		internalLog("error", "GetAssetLastTelmetries-Unmarshal: ")
		internalLog("error", err)
		return []Telmetry{}, err
	}

	return tagTelemetries.Result.Telmetries, nil
}

func DeleteAsset(assetID string, privateToken string, configs ...CervelloConfigurations) error {
	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return err
	}
	//https://api.release.cervello.io/application/v1/organizations/ffae83db-f249-4530-8db3-7298fdb75777/applications/a2d8a017-cf08-4c9d-a01c-6dec05af45fd/assets/59b0fee1-a0d3-4d33-bdf3-27864a0f2745
	resource := fmt.Sprintf(
		"/application/v1/organizations/%s/applications/%s/assets/%s",
		config.OrganizationID,
		config.ApplicationID,
		assetID,
	)

	var responseBodyObject baseResponse

	response, err := makeHTTPRequest("DELETE", htppRequest{
		Resource: resource,
		MetaData: "DeleteAsset",
		Token:    privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return err
	}
	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "DeleteAsset unmarshal response", err)
		return err
	}
	if responseBodyObject.Code != 2004 {
		return errors.New("Failed to DeleteAsset " + string(response.ResponseBody))
	}
	return nil

}

//GetAssetFiltered ...
//
// to get Assets by refrenceID set filter referenceName eq 43224-dfdfs-4323432-dsfsfdfdf
func GetAssetFiltered(queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) (*Asset, error) {
	// GET https://api.staging.cervello.io/application/v1/organizations/0feb9329-cf80-4c61-812b-c67eb74d5c33/applications/de1be3f0-6b43-47e8-8c23-984b07b0bf51/assets/
	// ?pageSize=1&filters[0][key]=referenceName&filters[0][operator]=eq&filters[0][value]=2388sa
	// headers [Authorization: "Bearer {token}"]

	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return nil, err
	}

	resource := fmt.Sprintf(
		"/application/v1/organizations/%s/applications/%s/assets",
		config.OrganizationID,
		config.ApplicationID,
	)

	var responseBodyObject getAssetResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource:    resource,
		MetaData:    "GetAssetFiltered",
		QueryParams: queryParam,
		Token:       privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetAssetFiltered unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "GetAssetFiltered error", string(response.ResponseBody))
	}

	if responseBodyObject.Result.Assets == nil || len(responseBodyObject.Result.Assets) == 0 {
		return &Asset{}, nil
	}
	return &responseBodyObject.Result.Assets[0], nil
}
func GetAssetsParentAssets(assetID string, queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) ([]AssetParentAsset, error) {
	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return nil, err
	}

	//https://api.release.cervello.io/application/v1/organizations/e4550d32-48f7-4f9c-bdf9-95df8f788140/applications/5dac9016-e1ef-4946-beb1-7f289b0977ca/assets/327f7e56-5555-4aeb-91ca-20559f2dd618/parents
	resource := fmt.Sprintf(
		"/application/v1/organizations/%s/applications/%s/assets/%s/parents",
		config.OrganizationID,
		config.ApplicationID,
		assetID,
	)

	var responseBodyObject getAssetParentAssetsResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource:    resource,
		MetaData:    "GetAssetsParentAssets",
		QueryParams: queryParam,
		Token:       privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetAssetsParentAssets unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "GetAssetsParentAssets error", string(response.ResponseBody))
	}

	return responseBodyObject.Result, nil

}

func GetAssetsChildrenAssets(assetID string, queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) ([]ChildAsset, error) {
	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return nil, err
	}
	//https://api.release.cervello.io/application/v1/organizations/:organistationId/applications/:applicationId/assets/:assetId/children
	resource := fmt.Sprintf(
		"/application/v1/organizations/%s/applications/%s/assets/%s/children",
		config.OrganizationID,
		config.ApplicationID,
		assetID,
	)

	var responseBodyObject getChildAssetResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource:    resource,
		MetaData:    "GetAssetsChildrenAssets",
		QueryParams: queryParam,
		Token:       privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetAssetsChildrenAssets unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "GetAssetsChildrenAssets error", string(response.ResponseBody))
	}

	return responseBodyObject.Result.Assets, nil
}
func GetAssetAlarms(assetID string, queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) ([]Alarm, error) {
	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return nil, err
	}

	//http://api.demo.cervello.io/data/v1/organizations/6b099687-0ae0-4df0-9273-a06bd146eb1f/applications/75050789-dd8f-4cfe-9b86-0136376be621/assets/6b66a1b4-6f8c-431b-99e5-812ba9bf0106/alarms
	resource := fmt.Sprintf(
		"/data/v1/organizations/%s/applications/%s/assets/%s/alarms",
		config.OrganizationID,
		config.ApplicationID,
		assetID,
	)

	var responseBodyObject getAlarmsResponse
	response, err := makeHTTPRequest("GET", htppRequest{
		Resource:    resource,
		MetaData:    "GetAssetAlarms",
		QueryParams: queryParam,
		Token:       privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}
	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetAssetAlarms unmarshal response", err)
		return nil, err
	}
	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "GetAssetAlarms error", string(response.ResponseBody))
	}
	return responseBodyObject.Result.Alarms, nil
}

func GetAssetAttributesFiltered(assetID string, queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) (map[string]interface{}, error) {
	// GET https://api.staging.cervello.io/data/v1/organizations/9a450c80-8511-4de5-8a1c-a3ecc6e9c862/applications/535880fd-91a5-4c2f-b7f6-4fe21869c730/assets/c428a94e-30bd-44e3-9970-b4c1ee558321/attributes
	// ?pageSize=100
	// headers [Authorization: "Bearer {token}"]

	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return nil, err
	}

	resource := fmt.Sprintf(
		"/data/v1/organizations/%s/applications/%s/assets/%s/attributes",
		config.OrganizationID,
		config.ApplicationID,
		assetID,
	)

	responseBodyObject := new(map[string]interface{})

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource:    resource,
		MetaData:    "GetAssetAttributesFiltered",
		QueryParams: queryParam,
		Token:       privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetAssetAttributesFiltered unmarshal response", err)
		return nil, err
	}

	return *responseBodyObject, nil
}
