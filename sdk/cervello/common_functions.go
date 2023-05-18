package cervello

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// Pagination ...
type Pagination struct {
	PageNumber int
	PageSize   int
}

// Filter ...
type Filter struct {
	Key   string
	Op    string
	Value string
}

type QueryParams struct {
	Filters       []Filter
	PaginationObj Pagination
	Custom        map[string]interface{}
}

// GetOrgID ...

// GetOrgID ...
func GetOrgID() string {
	return envOrganizationID
}

// GetAppID ...
func GetAppID() string {
	return envApplicationID
}

// GetApiUrl ...
func GetApiUrl() string {
	return envAPIURL
}

func addFiltersToRequest(data url.Values, filters []Filter) url.Values {
	for i, filter := range filters {
		data.Set(fmt.Sprintf("filters[%d][key]", i), filter.Key)
		data.Set(fmt.Sprintf("filters[%d][operator]", i), filter.Op)
		data.Set(fmt.Sprintf("filters[%d][value]", i), filter.Value)
	}

	return data
}

func addQueryParamsToRequest(data url.Values, queryParam QueryParams) url.Values {
	pageNumber := "1"
	pageSize := "10"
	// add filters
	for i, filter := range queryParam.Filters {
		data.Set(fmt.Sprintf("filters[%d][key]", i), filter.Key)
		data.Set(fmt.Sprintf("filters[%d][operator]", i), filter.Op)
		data.Set(fmt.Sprintf("filters[%d][value]", i), filter.Value)
	}
	// add pagination
	if queryParam.PaginationObj.PageNumber > 0 {
		pageNumber = strconv.Itoa(queryParam.PaginationObj.PageNumber)
	}
	if queryParam.PaginationObj.PageSize > 0 {
		pageSize = strconv.Itoa(queryParam.PaginationObj.PageSize)
	}
	data.Set("pageNumber", pageNumber)
	data.Set("pageSize", pageSize)

	// add custom queryParams
	for key, value := range queryParam.Custom {
		str := fmt.Sprintf("%v", value)
		data.Set(key, str)
	}
	return data
}

func addAuthHeader(req *http.Request, privateToken string) {
	if privateToken == "" {
		req.Header.Add("Authorization", "Bearer "+token.AccessToken)

	} else {
		req.Header.Add("Authorization", "Bearer "+privateToken)
	}
}

func parseCervelloConfig(isAppMandatory bool, configs []CervelloConfigurations) (*CervelloConfigurations, error) {
	config := &CervelloConfigurations{
		OrganizationID: envOrganizationID,
		ApplicationID:  envApplicationID,
	}
	if len(configs) > 0 {
		if configs[0].OrganizationID == "" {
			return nil, errors.New("organization ID is required")
		}
		config.OrganizationID = configs[0].OrganizationID
		if configs[0].ApplicationID == "" {
			if isAppMandatory {
				return nil, errors.New("application ID is required")
			}
			return config, nil
		}
		config.ApplicationID = configs[0].ApplicationID
	}
	return config, nil
}
