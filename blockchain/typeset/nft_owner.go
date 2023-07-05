package typeset

import (
	"time"

	"github.com/consolelabs/mochi-typeset/common/address/typeset"
)

// NftToken: nft token information
type NftOwner struct {
	// Id: internal id
	Id int64 `json:"id"`
	// OwnerAddress: address of the owner
	OwnerAddress *typeset.AccountAddress `json:"owner_address,omitempty"`
	// CollectionAddress: address of the collection contract
	CollectionAddress string `json:"collection_address,omitempty"`
	// TokenId: token id of the NFT
	TokenId string `json:"token_id,omitempty"`
	// ChainId: chain id of the contract, using internal system
	ChainId int64 `json:"chain_id,omitempty"`
	// Amount: amount of the NFT (erc721 = 1 by default)
	Amount int64 `json:"amount,omitempty"`

	// <-:false    Disable write permission
	Token           *NftToken  `json:"token" gorm:"foreignkey:TokenId,CollectionAddress;references:TokenId,CollectionAddress;<-:false"`
	CreatedTime     *time.Time `json:"created_time,omitempty"`
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
}
