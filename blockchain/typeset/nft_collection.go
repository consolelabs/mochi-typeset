package typeset

import "time"

type NftCollection struct {
	// Id: internal id
	ID int64 `json:"id"`
	// Address: address of the collection contract
	Address string `json:"collection_address"`
	// Name: name of the collection contract
	Name string `json:"name"`
	// Symbol: symbol of the collection contract
	Symbol string `json:"symbol"`
	// ChainId: chain id of the collection contract
	ChainId int64 `json:"chain_id"`
	// ERCFormat: format token of the collection contract
	ERCFormat string `json:"erc_format"`
	// Supply: supply of the collection contract
	Supply uint64 `json:"supply"`
	// Image: image of the collection contract
	Image string `json:"image"`
	// Description: description of the collection contract
	Description string `json:"description"`
	// ContractScan: contract scan of the collection contract
	ContractScan string `json:"contract_scan"`
	// Discord: discord of the collection contract
	Discord string `json:"discord"`
	// Twitter: twitter of the collection contract
	Twitter string `json:"twitter"`
	// Website: website of the collection contract
	Website         string    `json:"website"`
	CreatedTime     time.Time `json:"created_time,omitempty"`
	LastUpdatedTime time.Time `json:"last_updated_time,omitempty"`
}
