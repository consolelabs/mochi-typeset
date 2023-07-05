package typeset

// NftTokenAttribute: nft token attribute
type NftTokenAttribute struct {
	// CollectionAddress: address of the collection contract
	CollectionAddress string `json:"collection_address"`
	// TokenId: token onchain id of the NFT
	TokenId string `json:"token_id"`

	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
	Count     uint64 `json:"count"`
	Rarity    string `json:"rarity"`
	Frequency string `json:"frequency"`
}

// NftTokenRarity: nft token rarity
type NftTokenRarity struct {
	Rank   uint64 `json:"rank"`
	Score  string `json:"score"`
	Total  uint64 `json:"total"`
	Rarity string `json:"rarity,omitempty"`
}

// NftTokenAttributeCount: nft token attribute count
type NftTokenAttributeCount struct {
	Id    int64   `json:"id"`
	Count int64   `json:"count"`
	Ratio float64 `json:"ratio"`
}
