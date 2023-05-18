package cervello

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
)

// Alarm .....
type Alarm struct {
	ID              string                 `json:"id,omitempty"`
	OrganizationID  string                 `json:"organizationId,omitempty"`
	ApplicationID   string                 `json:"applicationId,omitempty"`
	OriginatorID    string                 `json:"originatorId,omitempty"`
	OrginatorName   string                 `json:"originatorName,omitempty"`
	StartTime       string                 `json:"startTime,omitempty"`
	Type            string                 `json:"type"`
	ActiveTime      string                 `json:"activeTime,omitempty"`
	ActiveUserID    string                 `json:"activeUserId,omitempty"`
	ActiveUserEmail string                 `json:"activeUserEmail,omitempty"`
	ActiveReason    string                 `json:"activeReason,omitempty"`
	AckTime         string                 `json:"ackTime,omitempty"`
	AckUserID       string                 `json:"ackUserId,omitempty"`
	AckUserEmail    string                 `json:"ackUserEmail,omitempty"`
	AckReason       string                 `json:"ackReason,omitempty"`
	EndTime         string                 `json:"endTime,omitempty"`
	ClearTime       string                 `json:"clearTime,omitempty"`
	ClearUserID     string                 `json:"clearUserId,omitempty"`
	ClearUserEmail  string                 `json:"clearUserEmail,omitempty"`
	ClearReason     string                 `json:"clearReason,omitempty"`
	Details         string                 `json:"details,omitempty"`
	Severity        int                    `json:"severity"`
	Status          int                    `json:"status"`
	Title           string                 `json:"title"`
	CreatedAt       string                 `json:"createdAt,omitempty"`
	UpdatedAt       string                 `json:"updatedAt,omitempty"`
	ResourceName    string                 `json:"resourceName,omitempty"`
	IsNew           bool                   `json:"isNew,omitempty"`
	Tags            []string               `json:"tags,omitempty"`
	CustomFields    map[string]interface{} `json:"customFields,omitempty"`
	MetaData        map[string]interface{} `json:"metadata,omitempty"`
	Action          string                 `json:"action,omitempty"`

	// verfication user_id
}

// BulkAlarmsActionsParams .....
type BulkAlarmsActionsParams struct {
	DeviceId       string `json:"deviceId,omitempty"`
	AlarmID        string `json:"alarmId,omitempty"`
	Action         string `json:"action,omitempty"`
	OrganizationId string `json:"organizationId,omitempty"`
	ApplicationId  string `json:"applicationId,omitempty"`
	Topic          string `json:"topic,omitempty"`
}

//GetAlarmByID ....
func GetAlarmByID(deviceID string, alarmID string, privateToken string, configs ...CervelloConfigurations) (*Alarm, error) {
	// TODO : implement get Alarm by ID
	// TODO : change alarm data model after adding Tags
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}

	//https://api.release.cervello.io/data/v1/organizations/e7738511-2e82-4a2a-b965-4460cd80a2fb/devices/b744e761-e5f8-40af-86d9-e16d31dac269/alarms
	resource := fmt.Sprintf(
		"/data/v1/organizations/%s/devices/%s/alarms",
		config.OrganizationID,
		deviceID,
	)

	var responseBodyObject *getAlarmsResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource: resource,
		MetaData: "GetAlarmByID",
		QueryParams: QueryParams{
			Filters: []Filter{
				{Key: "id", Op: "eq", Value: alarmID},
			},
		},
		Token: privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetAlarmByID unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "GetAlarmByID error", string(response.ResponseBody))
	}

	return &responseBodyObject.Result.Alarms[0], nil

}

//ChangeAlarmStatus ....
//
// used to change alarm status , the status must be on of [1,2,3,4]
func ChangeAlarmStatus(deviceID string, alarmID string, status int, privateToken string, configs ...CervelloConfigurations) error {
	resource := ""
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return err
	}

	switch status {
	case 1:
		resource = fmt.Sprintf(
			"/compose/v1/organizations/%s/alarms/%s/verify",
			config.OrganizationID,
			alarmID,
		)
		break
	case 2:
		resource = fmt.Sprintf(
			"/compose/v1/organizations/%s/alarms/%s/acknowledge",
			config.OrganizationID,
			alarmID,
		)
		break
	case 3:
		resource = fmt.Sprintf(
			"/compose/v1/organizations/%s/alarms/%s/clear",
			config.OrganizationID,
			alarmID,
		)

		break
	case 4:
		resource = fmt.Sprintf(
			"/compose/v1/organizations/%s/alarms/%s/clear",
			config.OrganizationID,
			alarmID,
		)
		break
	default:
		return errors.New("the status must be on of [0,1,2,3,4]")
	}

	var responseBodyObject baseResponse
	response, err := makeHTTPRequest("PUT", htppRequest{
		Resource: resource,
		MetaData: "ChangeAlarmStatus",
		Token:    privateToken,
	})

	if err != nil {
		internalLog("error", err)
		return err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "ChangeAlarmStatus unmarshal response", err)
		return err
	}

	if !strings.Contains(responseBodyObject.MessageKeys, "success") {
		return errors.New("error while changing alarm status " + string(response.ResponseBody))
	}
	return nil
}

//GetDeviceAlarmFiltered ....
// example get Active Alarms
//  filters.key = status
// 	filters.Op = eq
//  filters.Value = 1
func GetDeviceAlarmFiltered(deviceID string, queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) ([]Alarm, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}
	//https://api.release.cervello.io/data/v1/organizations/e4550d32-48f7-4f9c-bdf9-95df8f788140/devices/b744e761-e5f8-40af-86d9-e16d31dac269/alarms?filters[0][key]=status&filters[0][operator]=eq&filters[0][value]=3
	resource := fmt.Sprintf(
		"/data/v1/organizations/%s/devices/%s/alarms",
		config.OrganizationID,
		deviceID,
	)

	var responseBodyObject *getAlarmsResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource:    resource,
		MetaData:    "GetDeviceAlarmFiltered",
		QueryParams: queryParam,
		Token:       privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetDeviceAlarmFiltered unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "GetDeviceAlarmFiltered error", string(response.ResponseBody))
	}

	return responseBodyObject.Result.Alarms, nil
}

//CreateDeviceAlarm .....
// originator ID is required inisde the passed alarm object
func CreateDeviceAlarm(originatordID string, alarm Alarm, privateToken string, configs ...CervelloConfigurations) (Alarm, error) {
	//http://api.cervello.local/compose/v1/organizations/e4a58a52-6f9c-4b90-a004-e1f5ac9d5208/applications/b25b6627-8793-4601-8e41-ac2b893168a1/devices/ab43de18-9882-443d-9cf2-30700a068b5f/alarms
	//https: //api.release.cervello.io/compose/v1/organizations/e4550d32-48f7-4f9c-bdf9-95df8f788140/applications/5dac9016-e1ef-4946-beb1-7f289b0977ca/devices/d12d4369-5718-4ade-8fab-d74184e6d20e/alarms

	config, err := parseCervelloConfig(true, configs)
	if err != nil {
		return Alarm{}, err
	}

	resource := fmt.Sprintf(
		"/compose/v1/organizations/%s/applications/%s/devices/%s/alarms",
		config.OrganizationID,
		config.ApplicationID,
		originatordID,
	)

	var responseBodyObject createAlarmResponse

	response, err := makeHTTPRequest("POST", htppRequest{
		Resource: resource,
		MetaData: "CreateDeviceAlarm",
		Token:    privateToken,
		Body:     alarm,
	})

	if err != nil {
		internalLog("error", err)
		return Alarm{}, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "CreateDeviceAlarm unmarshal response", err)
		return Alarm{}, err
	}
	if responseBodyObject.Code != 2001 {
		return Alarm{}, errors.New("Failed to CreateDeviceAlarm: " + string(response.ResponseBody))
	}

	return responseBodyObject.Result, nil
}

//ChangeAlarmStatusBulk ....
//
// used to change alarm status bulk alarms , the status must be on of [1,2,3,4]
func ChangeAlarmStatusBulk(alarms []BulkAlarmsActionsParams, privateToken string) error {
	for _, alarmParams := range alarms {
		alarmParams.OrganizationId = GetOrgID()
		alarmParams.ApplicationId = GetOrgID()
		alarmParams.Topic = fmt.Sprintf("/cervello/devices/%v/alarm", alarmParams.DeviceId)
		messageBytes, err := json.Marshal(alarmParams)
		if err != nil {
			return err
		}

		log.Printf(string(messageBytes))
		if err := CervelloNatsConnection.Publish(NatsDeviceDataMessageTopic, messageBytes); err != nil {
			return err
		}
	}
	return nil
}

//GetOrgAlarmsFiltered ....
func GetAlarmsFiltered(queryParam QueryParams, privateToken string, configs ...CervelloConfigurations) (*[]Alarm, error) {
	config, err := parseCervelloConfig(false, configs)
	if err != nil {
		return nil, err
	}

	//http://api.demo.cervello.io/compose/v1/organizations/6b099687-0ae0-4df0-9273-a06bd146eb1f/alarms
	resource := fmt.Sprintf(
		"/data/v1/organizations/%s/alarms",
		config.OrganizationID,
	)
	var responseBodyObject *getAlarmsResponse

	response, err := makeHTTPRequest("GET", htppRequest{
		Resource:    resource,
		MetaData:    "GetOrgAlarmsFiltered",
		QueryParams: queryParam,
		Token:       privateToken,
	})
	if err != nil {
		internalLog("error", err)
		return nil, err
	}

	if err = json.Unmarshal(response.ResponseBody, &responseBodyObject); err != nil {
		internalLog("error", "GetOrgAlarmsFiltered unmarshal response", err)
		return nil, err
	}

	if responseBodyObject.Code != 2003 {
		return nil, fmt.Errorf("%v-%v", "GetOrgAlarmsFiltered error", string(response.ResponseBody))
	}

	return &responseBodyObject.Result.Alarms, nil

}
