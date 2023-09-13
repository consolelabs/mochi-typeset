package typeset

const (
	CHAIN_TYPE_EVM     = "evm"
	CHAIN_TYPE_SOLANA  = "solana"
	CHAIN_TYPE_RONIN   = "ronin"
	CHAIN_TYPE_BITCOIN = "bitcoin"
	CHAIN_TYPE_TON     = "ton"
)

func IsEvm(chainType string) bool {
	return chainType == CHAIN_TYPE_EVM
}

func IsSolana(chainType string) bool {
	return chainType == CHAIN_TYPE_SOLANA
}

func IsRonin(chainType string) bool {
	return chainType == CHAIN_TYPE_RONIN
}

func IsBitcoin(chainType string) bool {
	return chainType == CHAIN_TYPE_BITCOIN
}

func IsTon(chainType string) bool {
	return chainType == CHAIN_TYPE_TON
}
