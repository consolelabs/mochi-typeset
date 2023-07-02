package typeset

type OnchainInstructuion struct {
	ContractStructure string      `json:"contract_structure"`
	ContractAddress   string      `json:"contract_address"`
	FunctionName      string      `json:"function_name"`
	Params            []string    `json:"params"`
	ChainId           string      `json:"chain_id"`
	Metadata          interface{} `json:"metadata"`
}
