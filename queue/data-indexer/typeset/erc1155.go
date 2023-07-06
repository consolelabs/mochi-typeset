package typeset

import (
	"encoding/json"

	btypeset "github.com/consolelabs/mochi-typeset/blockchain/typeset"
	htypset "github.com/consolelabs/mochi-typeset/common/contract/hash/typeset"
	queue "github.com/consolelabs/mochi-typeset/queue"
)

// TransferSingleEventMessage is a message for transfer event message NFT
type TransferSingleEventMessage btypeset.EvmEvent

func ToTransferSingleEventMessage(msg *queue.KafkaMessage) (event *TransferSingleEventMessage) {
	if msg.Type != htypset.TRANSFER_SINGLE_HASH {
		return nil
	}

	err := json.Unmarshal(msg.Data, &event)
	if err != nil {
		return nil
	}

	return event
}
