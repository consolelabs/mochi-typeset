package typeset

import (
	"encoding/json"

	"github.com/consolelabs/mochi-typeset/common/address/typeset"
	queue "github.com/consolelabs/mochi-typeset/queue"
)

// NftIndexMessage is a message for indexing NFT
type NftIndexMessage struct {
	// Type: indicate which logic consumer should use
	Type string `json:"type"`
	// ContractAddress: address of the contract
	ContractAddress string `json:"contract_address"`
	// ChainId: chain id of the contract
	ChainId int64 `json:"chain_id"`
	// From: address of the sender
	From *typeset.AccountAddress `json:"from"`
	// To: address of the receiver
	To *typeset.AccountAddress `json:"to"`
	// TokenId: token id of the NFT
	TokenId string `json:"token_id"`
	// Amount: amount of the NFT
	Amount int64 `json:"amount"`
	// EventTime: time of event
	EventTime int64 `json:"event_time"`
	// TransactionHash: hash of the transaction
	TransactionHash string `json:"transaction_hash"`
	// BlockNumber: block number
	BlockNumber int64 `json:"block_number"`
}

func ToNftMessage(msg *queue.KafkaMessage) (nft *NftIndexMessage) {
	if msg.Type != "nft" {
		return nil
	}

	err := json.Unmarshal(msg.Data, &nft)
	if err != nil {
		return nil
	}

	return nft
}
