package typeset

import (
	blockchain "github.com/consolelabs/mochi-typeset/blockchain/typeset"
	typeset "github.com/consolelabs/mochi-typeset/typeset"
)

// NotifierMessage is the message that will be sent to the notifier queue
// TODO: This is a clone of mochi-notifications.NotifierMessage, need plan to move completely to typeset
type NotifierMessage struct {
	Type                   typeset.NotificationType `json:"type"`
	WalletTrackingMetadata *WalletTrackingMetadata  `json:"wallet_tracking_metadata"`
}

// WalletTrackingMetadata is the metadata for wallet tracking notification
type WalletTrackingMetadata struct {
	ProfileId     string                     `json:"profile_id"`
	ChainType     string                     `json:"chain_type"`
	WalletAddress string                     `json:"wallet_address"`
	WalletAlias   string                     `json:"wallet_alias"`
	Amount        typeset.TokenAmount        `json:"amount"`
	Token         typeset.Token              `json:"token"`
	Transaction   *blockchain.EvmTransaction `json:"transaction"`
}
