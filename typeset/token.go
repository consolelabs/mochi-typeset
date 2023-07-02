package typeset

type Token struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	Address  string `json:"address"`
	ChainId  string `json:"chain_id"`
}

type TokenAmount struct {
	AmountInToken string `json:"amount_in_token"`
	AmountInUsd   string `json:"amount_in_usd"`
	Decimals      int    `json:"decimals"`
}
