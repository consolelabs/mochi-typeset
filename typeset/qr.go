package typeset

type QrType int

const (
	QR_TYPE_AIRDROP QrType = iota
)

type Qr struct {
	// Type: Type of the QR code
	Type int         `json:"type"`
	Data interface{} `json:"data"`
}
