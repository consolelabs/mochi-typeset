package typeset

import (
	"time"
)

// Contract: contract information
type Contract struct {
	ID               int64      `json:"id"`
	LastUpdatedTime  *time.Time `json:"last_updated_time"`
	LastUpdatedBlock int64      `json:"last_updated_block"`
	CreationBlock    int64      `json:"creation_block"`
	CreatedTime      *time.Time `json:"created_time"`

	Address string `json:"address"`
	ChainId int64  `json:"chain_id"`
	Type    string `json:"type"`
}
