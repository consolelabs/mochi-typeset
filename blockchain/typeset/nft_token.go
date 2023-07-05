package typeset

import "time"

type NftToken struct {
	TokenId           string `json:"token_id"`
	CollectionAddress string `json:"collection_address,omitempty"`
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
	Amount            string `json:"amount,omitempty"`
	Image             string `json:"image,omitempty"`
	ImageCDN          string `json:"image_cdn,omitempty"`
	ThumbnailCDN      string `json:"thumbnail_cdn,omitempty"`
	ImageContentType  string `json:"image_content_type,omitempty"`
	RarityRank        uint64 `json:"rarity_rank,omitempty"`
	RarityScore       string `json:"rarity_score,omitempty"`

	Attributes      []NftTokenAttribute `json:"attributes" gorm:"-"` // disable write
	CreatedTime     *time.Time          `json:"created_time,omitempty"`
	LastUpdatedTime *time.Time          `json:"last_updated_time,omitempty"`
}
