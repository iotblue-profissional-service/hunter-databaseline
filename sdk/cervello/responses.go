package cervello

import "time"

type baseResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys,omitempty"`
	MessageKey  string `json:"messageKey,omitempty"`
	Message     string `json:"message,omitempty"`
}

type postAssetResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys,omitempty"`
	MessageKey  string `json:"messageKey,omitempty"`
	Message     string `json:"message,omitempty"`
	Result      Asset  `json:"result,omitempty"`
}

type getAssetResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys,omitempty"`
	MessageKey  string `json:"messageKey,omitempty"`
	Message     string `json:"message,omitempty"`
	Result      struct {
		Assets     []Asset     `json:"data"`
		TotalCount interface{} `json:"totalCount,omitempty"`
	} `json:"result"`
}

type getChildAssetResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys,omitempty"`
	MessageKey  string `json:"messageKey,omitempty"`
	Message     string `json:"message,omitempty"`
	Result      struct {
		Assets     []ChildAsset `json:"data"`
		TotalCount interface{}  `json:"totalCount,omitempty"`
	} `json:"result"`
}

type getAssetByIDResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys,omitempty"`
	MessageKey  string `json:"messageKey,omitempty"`
	Message     string `json:"message,omitempty"`
	Result      Asset  `json:"result"`
}

type getAssetParentAssetsResponse struct {
	Code        int                `json:"code"`
	MessageKeys string             `json:"messageKeys,omitempty"`
	MessageKey  string             `json:"messageKey,omitempty"`
	Message     string             `json:"message,omitempty"`
	Result      []AssetParentAsset `json:"result"`
}

type getAlarmsResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys,omitempty"`
	MessageKey  string `json:"messageKey,omitempty"`
	Message     string `json:"message,omitempty"`
	Result      struct {
		Alarms     []Alarm `json:"data"`
		TotalCount int     `json:"totalCount"`
	} `json:"result"`
}

type getDeviceAssetsResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys,omitempty"`
	MessageKey  string `json:"messageKey,omitempty"`
	Message     string `json:"message,omitempty"`
	Result      struct {
		TotalCount   int           `json:"totalCount"`
		DeviceAssets []DeviceAsset `json:"data"`
	} `json:"result"`
}

type getAssetDevicesResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys,omitempty"`
	MessageKey  string `json:"messageKey,omitempty"`
	Message     string `json:"message,omitempty"`
	Result      struct {
		AssetDevices []Device `json:"data"`
	} `json:"result"`
}

type getDeviceResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys,omitempty"`
	MessageKey  string `json:"messageKey,omitempty"`
	Message     string `json:"message,omitempty"`
	Result      struct {
		Devices    []Device `json:"data"`
		TotalCount int      `json:"totalCount"`
	} `json:"result"`
}

type getDeviceBYIDResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys,omitempty"`
	MessageKey  string `json:"messageKey,omitempty"`
	Message     string `json:"message,omitempty"`
	Result      Device `json:"result"`
}

type getTagTelemtriesResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys,omitempty"`
	MessageKey  string `json:"messageKey,omitempty"`
	Message     string `json:"message,omitempty"`
	Result      struct {
		Telmetries []Telmetry `json:"data"`
	} `json:"result"`
}
type postDeviceResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys,omitempty"`
	MessageKey  string `json:"messageKey,omitempty"`
	Message     string `json:"message,omitempty"`
	Result      struct {
		Device Device `json:"device"`
	} `json:"result,omitempty"`
	//Result      map[string]Device `json:"result,omitempty"`
}

type createAlarmResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys,omitempty"`
	MessageKey  string `json:"messageKey,omitempty"`
	Message     string `json:"message,omitempty"`
	Result      Alarm  `json:"result"`
}

type deviceCredentialsResponse struct {
	Code        int               `json:"code"`
	MessageKeys string            `json:"messageKeys,omitempty"`
	MessageKey  string            `json:"messageKey,omitempty"`
	Message     string            `json:"message,omitempty"`
	Result      DeviceCredentials `json:"result,omitempty"`
}
type getModbusConfigurationResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys"`
	Message     string `json:"message"`
	Result      struct {
		OrganizationID string                      `json:"organizationId"`
		DeviceID       string                      `json:"deviceId"`
		Configuration  []ModbusConfiguration       `json:"configuration"`
		Schedule       ModbusConfigurationSchedule `json:"schedule"`
		ID             string                      `json:"id"`
		CreatedAt      time.Time                   `json:"createdAt"`
		UpdatedAt      time.Time                   `json:"updatedAt"`
	} `json:"result"`
}

type getBacnetConfigurationResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys"`
	Message     string `json:"message"`
	Result      struct {
		ID             string `json:"id"`
		DeviceID       string `json:"deviceId"`
		OrganizationID string `json:"organizationId"`
		Configuration  struct {
			Objects []BacnetConfiguration `json:"objects"`
		} `json:"configuration"`
		Schedule  BacnetConfigurationSchedule `json:"schedule"`
		CreatedAt time.Time                   `json:"createdAt"`
		UpdatedAt time.Time                   `json:"updatedAt"`
	} `json:"result"`
}

type getDeviceCommandsResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys"`
	Message     string `json:"message"`
	Result      struct {
		Commands   []Command   `json:"data"`
		TotalCount int         `json:"totalCount"`
		PageNumber interface{} `json:"pageNumber"`
		PageSize   interface{} `json:"pageSize"`
	} `json:"result"`
}

type getCommandTemplatesResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys"`
	Message     string `json:"message"`
	Result      struct {
		CommandTemplate []CommandTemplate `json:"data"`
		TotalCount      int               `json:"totalCount"`
		PageNumber      interface{}       `json:"pageNumber"`
		PageSize        interface{}       `json:"pageSize"`
	} `json:"result"`
}

type createCommandTemplatesResponse struct {
	Code        int             `json:"code"`
	MessageKeys string          `json:"messageKeys"`
	Message     string          `json:"message"`
	Result      CommandTemplate `json:"result"`
}

type createDeviceCommandResponse struct {
	Code        int     `json:"code"`
	MessageKeys string  `json:"messageKeys"`
	Message     string  `json:"message"`
	Result      Command `json:"result"`
}

type createOrganizationCredentialResponse struct {
	Code        int                    `json:"code"`
	MessageKeys string                 `json:"messageKeys"`
	Message     string                 `json:"message"`
	Result      OrganizationCredential `json:"result"`
}

type getOrganizationCredentialResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys"`
	Message     string `json:"message"`
	Result      struct {
		Credentails []OrganizationCredential `json:"data"`
	} `json:"result"`
}

type getMultiCastGroupDevicesResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys,omitempty"`
	MessageKey  string `json:"messageKey,omitempty"`
	Message     string `json:"message,omitempty"`
	Result      struct {
		Devices []Device `json:"devices"`
	} `json:"result"`
}

type getApplicationVariablesResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys"`
	Message     string `json:"message"`
	Result      struct {
		ApplicationVariables []ApplicationVariable `json:"data"`
		TotalCount           int                   `json:"totalCount"`
		PageNumber           interface{}           `json:"pageNumber"`
		PageSize             interface{}           `json:"pageSize"`
	} `json:"result"`
}

type getApplicationConfigurationsResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys"`
	Message     string `json:"message"`
	Result      struct {
		ApplicationConfiguration []ApplicationConfiguration `json:"data"`
		TotalCount               int                        `json:"totalCount"`
		PageNumber               interface{}                `json:"pageNumber"`
		PageSize                 interface{}                `json:"pageSize"`
	} `json:"result"`
}

type getOrganisationsResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys"`
	Message     string `json:"message"`
	Result      struct {
		Data       []Organization `json:"data"`
		TotalCount int            `json:"totalCount"`
		PageNumber string         `json:"pageNumber"`
		PageSize   string         `json:"pageSize"`
	} `json:"result"`
}

type getOrganisationResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys"`
	Message     string `json:"message"`
	Result      struct {
		Data Organization `json:"data"`
	} `json:"result"`
}
