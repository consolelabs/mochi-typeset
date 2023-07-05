package typeset

import (
	"time"

	"github.com/consolelabs/mochi-typeset/common/address/typeset"
)

// NftTransfer: nft transfer information
type NftTransfer struct {
	// Id: internal id
	Id int64 `json:"id"`
	// From: address of the sender
	From *typeset.AccountAddress `json:"from"`
	// To: address of the receiver
	To *typeset.AccountAddress `json:"to"`
	// TokenId: token id of the NFT
	TokenId string `json:"token_id"`
	// EventTime: time of event
	EventTime int64 `json:"event_time"`
	// TransactionHash: hash of the transaction
	TransactionHash string `json:"transaction_hash"`
	// BlockNumber: block number
	BlockNumber int64 `json:"block_number"`
	// ContractAddress: address of the collection contract
	ContractAddress string    `json:"contract_address"`
	CreatedTime     time.Time `json:"created_time"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}
