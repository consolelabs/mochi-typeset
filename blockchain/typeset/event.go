package typeset

import (
	"encoding/json"

	"github.com/consolelabs/mochi-typeset/common/contract/hash/typeset"
	queue "github.com/consolelabs/mochi-typeset/queue"
)

// EvmEvent: event log type for evm
type EvmEvent struct {
	ChainId          int64    `json:"chain_id"`
	LogIndex         int64    `json:"log_index"`
	TransactionHash  string   `json:"transaction_hash"`
	TransactionIndex int64    `json:"transaction_index"`
	BlockHash        string   `json:"block_hash"`
	BlockNumber      int64    `json:"block_number"`
	Address          string   `json:"address"`
	Data             []byte   `json:"data"`
	Topics           []string `json:"topics"`
}

func ToEvmEvent(msg *queue.KafkaMessage) (event *EvmEvent) {
	if msg.Type != typeset.TRANSFER_SINGLE_HASH {
		return nil
	}

	err := json.Unmarshal(msg.Data, &event)
	if err != nil {
		return nil
	}

	return event
}
