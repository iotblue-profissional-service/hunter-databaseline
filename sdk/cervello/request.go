package cervello

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type htppRequest struct {
	Resource    string
	QueryParams QueryParams
	Body        interface{}
	MetaData    interface{}
	Token       string
}
type htppResponse struct {
	Response     *http.Response
	ResponseBody []byte
}

func makeHTTPRequest(requestType string, request htppRequest) (*htppResponse, error) {
	var requestBodyJSON []byte
	var urlStr string
	var err error
	data := url.Values{}
	u, _ := url.ParseRequestURI(envAPIURL)
	u.Path = request.Resource

	switch requestType {
	case "GET":
		data = addQueryParamsToRequest(data, request.QueryParams)
		u.RawQuery = data.Encode()
	case "DELETE":
		data = addQueryParamsToRequest(data, request.QueryParams)
		u.RawQuery = data.Encode()
	case "POST", "PUT":
		if request.Body != nil {
			requestBodyJSON, err = json.Marshal(request.Body)
			if err != nil {
				internalLog("error", request.MetaData, "marshal request body", err)
				return nil, err
			}
		}
	default:
		return nil, errors.New("un supported http verb")
	}
	urlStr = u.String()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(requestType, urlStr, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		internalLog("error", request.MetaData, "new request", err)
		return nil, err
	}

	addAuthHeader(req, request.Token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		internalLog("error", request.MetaData, "do request", err)
		return nil, err
	}

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		internalLog("error", request.MetaData, "read response body", err)
		return nil, err
	}

	resp.Body.Close()

	return &htppResponse{
		Response:     resp,
		ResponseBody: f,
	}, nil
}
