package cervello

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Telmetry struct {
	ResourceID   string      `json:"resourceId,omitempty"`
	Key          string      `json:"key,omitempty"`
	ValueType    string      `json:"valueType,omitempty"`
	Value        interface{} `json:"value,omitempty"`
	Time         string      `json:"time,omitempty"`
	ResourceName string      `json:"resourceName,omitempty"`
}

func CountTagDevices(tag string, privateToken string) (int, error) {
	resource := fmt.Sprintf(
		"/device/v1/organizations/%s/devices",
		envOrganizationID,
	)

	var responseBodyObject getDeviceResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource: resource,
		MetaData: "CountTagDevices",
		QueryParams: QueryParams{
			Filters: []Filter{
				{Key: "tags", Op: "contains", Value: tag},
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
		internalLog("error", "CountTagDevices unmarshal response", err)
		return 0, err
	}

	if responseBodyObject.Code != 2003 {
		return 0, fmt.Errorf("%v-%v", "CountTagDevices error", string(response.ResponseBody))
	}

	return responseBodyObject.Result.TotalCount, nil
}

//GetTagLastTelmetries ... to get all telemetries let key=""
func GetTagLastTelmetries(key string, tag string, queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) ([]Telmetry, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}
	//https://api.release.cervello.io/compose/v1/organizations/6df9aab1-0a16-4c24-9176-6a879a5a514e/tags/city-controller/telemetries?aggregation=last&filters[0][key]=key&filters[0][operator]=eq&filters[0][value]=key
	resource := fmt.Sprintf(
		"/compose/v1/organizations/%s/tags/%s/telemetries",
		config.OrganizationID,
		tag,
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
		internalLog("error", "GetTagLastTelmetries-NewRequest: ")
		internalLog("error", err)
		return []Telmetry{}, err
	}

	addAuthHeader(req, privateToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		internalLog("error", "GetTagLastTelmetries-Do: ")
		internalLog("error", err)
		return []Telmetry{}, err
	}

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		internalLog("error", "CountTagDevices-ReadAll: ")
		internalLog("error", err)
		return []Telmetry{}, err
	}

	resp.Body.Close()

	// Assets := new(getAssetResponse)
	var tagTelemetries getTagTelemtriesResponse

	err = json.Unmarshal(f, &tagTelemetries)
	if err != nil {
		internalLog("error", "GetTagLastTelmetries-Unmarshal: ")
		internalLog("error", err)
		return []Telmetry{}, err
	}

	return tagTelemetries.Result.Telmetries, nil
}
