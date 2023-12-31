package typeset

import "github.com/consolelabs/mochi-typeset/common/error/typeset"

type KafkaMessageType string

// TODO: update notification
const (
	KAFKA_MESSAGE_TYPE_AUDIT                     KafkaMessageType = "audit"
	KAFKA_MESSAGE_TYPE_TELEMETRY_NEW_TRANSACTION KafkaMessageType = "telemetry_new_transaction"
)

func (k KafkaMessageType) String() string {
	return string(k)
}

type KafkaMessage struct {
	Type   KafkaMessageType `json:"type"`
	Topic  string           `json:"topic"`
	Data   []byte           `json:"message"`
	Sender string           `json:"sender"`
}

// Validate validates the kafka message
// 1. Type should not be empty
// 2. Sender should not be empty
func (k *KafkaMessage) Validate() error {
	switch k.Type {
	case
		KAFKA_MESSAGE_TYPE_AUDIT,
		KAFKA_MESSAGE_TYPE_TELEMETRY_NEW_TRANSACTION:
		return nil
	default:
		return typeset.ERROR_FAIL_VALIDATION
	}
}
