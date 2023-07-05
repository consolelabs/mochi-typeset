package typeset

import "github.com/consolelabs/mochi-typeset/common/address/typeset"

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
}
