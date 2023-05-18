package cervello

const (
	// KafkaAlarmTopic ...
	KafkaAlarmTopic = "alarm"
	// KafkaAuditLogTopic ....
	KafkaAuditLogTopic = "audit_log"
	// KafkaSocketTopic ....
	KafkaSocketTopic = "socket"
	// NatsAlarmTopic ...
	NatsAlarmTopic = "devices.data.alarm"
	// NatsAuditLogTopic ....
	NatsAuditLogTopic = "user.auditlog"
	// NatsSocketTopic ....
	NatsSocketTopic = "notification.socket"
	// NatsSocketTopic ....
	NatsDeviceDataMessageTopic = "device.data.message"
	AlarmActionClear           = "clear_alarm"
	AlarmActionCreate          = "create_alarm"
	AlarmActionVerify          = "verify_alarm"
	AlarmActionAcknowldge      = "acknowledge_alarm"
	AlarmActionFalse           = "false_alarm"
)

//AlarmMessage ...
type AlarmMessage struct {
	OrganizationID string `json:"organizationId"`
	ApplicationID  string `json:"applicationId"`
	Alarm          Alarm  `json:"data"`
}

//AlarmActionsMessage ...
type AlarmActionsMessage struct {
	DeviceID     string `json:"deviceId"`
	OriginatorId string `json:"originatorId"`
	Action       string `json:"action"`
}

//SocketMessage ...
type SocketMessage struct {
	OrganizationID string            `json:"organizationId"`
	ResourceID     string            `json:"resourceId"`
	ResourceType   string            `json:"resourceType"`
	Data           map[string]string `json:"data"`
}

//AuditLogMessage ...
type AuditLogMessage struct {
	OrganizationID       string                 `json:"organizationId"`
	ApplicationID        string                 `json:"applicationId"`
	UserID               string                 `json:"userId"`
	UserEmail            string                 `json:"userEmail"`
	EntityID             string                 `json:"entityId"`
	Entity               string                 `json:"entity"`
	RequestServiceIssuer string                 `json:"requestServiceIssuer"`
	Action               string                 `json:"action"`
	Status               bool                   `json:"status"`
	Data                 map[string]interface{} `json:"data"`
}
