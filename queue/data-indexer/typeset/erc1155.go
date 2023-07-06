package typeset

import (
	"encoding/json"

	"github.com/consolelabs/mochi-typeset/common/contract/hash/typeset"
	queue "github.com/consolelabs/mochi-typeset/queue"
)

// TransferSingleEventMessage is a message for transfer event message NFT
type TransferSingleEventMessage struct {
	Type             string   `json:"type"`
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

func ToTransferSingleEventMessage(msg *queue.KafkaMessage) (event *TransferSingleEventMessage) {
	if msg.Type != typeset.TRANSFER_SINGLE_HASH {
		return nil
	}

	err := json.Unmarshal(msg.Data, &event)
	if err != nil {
		return nil
	}

	return event
}
