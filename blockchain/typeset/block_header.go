package typeset

// BlockHeader: small size version to check block
type BlockHeader struct {
	ChainId   int64  `json:"chain_id"`
	Number    int64  `json:"number"`
	Hash      string `json:"hash"`
	CreatedAt int64  `json:"created_at"`
}
