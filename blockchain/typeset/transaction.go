package typeset

// EvmTransaction: transaction type for evm
type EvmTransaction struct {
	ChainId          int64      `json:"chain_id"`
	Hash             string     `json:"hash"`
	Nonce            int64      `json:"nonce"`
	BlockHash        string     `json:"block_hash"`
	BlockNumber      int64      `json:"block_number"`
	TransactionIndex int64      `json:"transaction_index"`
	FromAddress      string     `json:"from_address"`
	ToAddress        string     `json:"to_address"`
	Value            string     `json:"value"`
	Gas              int64      `json:"gas"`
	GasPrice         int64      `json:"gas_price"`
	Input            []byte     `json:"input"`
	BlockTimestamp   int64      `json:"block_timestamp"`
	MaxFeePerGas     int64      `json:"max_fee_per_gas"`
	TransactionType  string     `json:"transaction_type"`
	Events           []EvmEvent `json:"events"`
}
