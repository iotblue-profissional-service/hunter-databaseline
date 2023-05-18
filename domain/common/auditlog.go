package common

import (
	"databaselineservice/sdk/cervello"
	"databaselineservice/sdk/logger"
	"encoding/json"
)

func PublishAuditLog(action string, entity string, entityId string, data interface{}) {
	// construct profile audit log message
	message := cervello.AuditLogMessage{
		Action:               action,
		ApplicationID:        cervello.GetAppID(),
		OrganizationID:       cervello.GetOrgID(),
		Entity:               entity,
		EntityID:             entityId,
		Status:               true,
		RequestServiceIssuer: "dataBaseLineService",
		UserEmail:            "super@cervello.local",
		Data: map[string]interface{}{
			"data": data,
		},
	}
	logger.LogMessage("info", "PushAuditLog - starts")

	// serialize the notification object
	auditlogSerializedMessage, _ := json.Marshal(message)
	// log.Printf("the socket message", string(socketMessage))
	// publish the serialized object
	// message in byte array is the second paramter
	// if err := cervello.CervelloNatsConnection.Publish(cervello.NatsAuditLogTopic, auditlogSerializedMessage); err != nil {
	// 	logger.LogMessage("error", "error nats publish", err)
	// }
	logger.LogMessage("info", "PushAuditLog - finished", auditlogSerializedMessage)
}
