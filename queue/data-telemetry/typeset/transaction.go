package typeset

import (
	"encoding/json"

	btype "github.com/consolelabs/mochi-typeset/blockchain/typeset"
	qtype "github.com/consolelabs/mochi-typeset/queue"
)

// TelemetryNewTransactionMessage is a message for new transaction message
type TelemetryNewTransactionMessage btype.EvmTransaction

// ToTelemetryNewTransactionMessage converts a KafkaMessage to a TelemetryNewTransactionMessage
func ToTelemetryNewTransactionMessage(msg *qtype.KafkaMessage) (event *TelemetryNewTransactionMessage) {
	if msg.Type != qtype.KAFKA_MESSAGE_TYPE_TELEMETRY_NEW_TRANSACTION {
		return nil
	}

	err := json.Unmarshal(msg.Data, &event)
	if err != nil {
		return nil
	}

	return event
}
