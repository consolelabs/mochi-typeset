package typeset

// https://ftmscan.com/address/0x9e3922FCBb28F40026dA2E58926f0f249d8A138d
// abi: https://ftmscan.com/address/0x9e3922FCBb28F40026dA2E58926f0f249d8A138d#code
// address: 0x9e3922FCBb28F40026dA2E58926f0f249d8A138d
// functionName: deposit
// params: ["10000000"]
// chainId: 250

// OnchainInstructions uses for the frontend to submit to the onchain contract
type OnchainInstructions struct {
	// ContractStructure: Abi / Idl
	ContractStructure string `json:"contract_structure"`
	// ContractAddress: Address of the contract
	ContractAddress string `json:"contract_address"`
	// FunctionName: Name of the function to call (in abi / idl json)
	FunctionName string `json:"function_name"`
	// Params: Parameters to pass to the function
	Params []string `json:"params"`
	// ChainId: Chain id to use (use within our database)
	ChainId int64 `json:"chain_id"`
	// Metadata: Metadata for the UI rendering
	Metadata interface{} `json:"metadata"`
}
