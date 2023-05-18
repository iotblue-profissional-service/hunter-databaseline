package cervello

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

// NatsConnection ...
type NatsConnection struct {
	Uri string
}

var conn *nats.Conn = nil
var (
	// CervelloNatsConnection  the connection used to publish and subscribe on kafka
	CervelloNatsConnection *NatsConnection
)

// json messages
// uri : 41.65.47.227:4222
func connect(uri string) *nats.Conn {
	if conn == nil {
		natsConnection, err := nats.Connect(uri)
		if err != nil {
			internalLog("error", "NATS connection error...")
			log.Fatal(err)
		}
		conn = natsConnection
	}
	return conn
}

// NewNatsConnection ...
func NewNatsConnection(URI string) *NatsConnection {
	CervelloNatsConnection = &NatsConnection{
		URI,
	}
	return CervelloNatsConnection
}

// Subscribe ...
func (conn *NatsConnection) Subscribe(topic string, cb func(message []byte)) {
	natsConnection := connect(conn.Uri)

	_, err := natsConnection.Subscribe(topic, func(msg *nats.Msg) {
		cb(msg.Data)
	})
	if err != nil {
		internalLog("error", fmt.Sprintf(`Message Distributor error while subscribing on topic %v`, topic))
		log.Fatal(err)
	}
}

// Publish ...
func (conn *NatsConnection) Publish(topic string, message []byte) error {
	natsConnection := connect(conn.Uri)
	return natsConnection.Publish(topic, message)
}
