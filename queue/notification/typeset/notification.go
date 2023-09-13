package typeset

import (
	blockchain "github.com/consolelabs/mochi-typeset/blockchain/typeset"
	friendtech "github.com/consolelabs/mochi-typeset/common/transaction/friendtech/typeset"
	typeset "github.com/consolelabs/mochi-typeset/typeset"
)

// NotifierMessage is the message that will be sent to the notifier queue
// TODO: This is a clone of mochi-notifications.NotifierMessage, need plan to move completely to typeset
type NotifierMessage struct {
	Type                      typeset.NotificationType   `json:"type"`
	WalletTrackingMetadata    *WalletTrackingMetadata    `json:"wallet_tracking_metadata"`
	KeyNewTransactionMetadata *KeyNewTransactionMetadata `json:"key_new_transaction_metadata"`
}

// WalletTrackingMetadata is the metadata for wallet tracking notification
type WalletTrackingMetadata struct {
	ProfileId     string                     `json:"profile_id"`
	ChainId       int64                      `json:"chain_id"`
	WalletAddress string                     `json:"wallet_address"`
	TxHash        string                     `json:"tx_hash"`
	WalletAlias   string                     `json:"wallet_alias"`
	IsSender      bool                       `json:"is_sender"`
	Amount        typeset.TokenAmount        `json:"amount"`
	Token         typeset.Token              `json:"token"`
	Transaction   *blockchain.EvmTransaction `json:"transaction"`
}

// KeyNewTransactionMetadata is the metadata for key new transaction notification
type KeyNewTransactionMetadata struct {
	ProfileId   string                               `json:"profile_id"`
	Transaction *friendtech.FriendTechKeyTransaction `json:"transaction"`
}
