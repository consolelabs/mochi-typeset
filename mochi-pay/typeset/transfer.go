package typeset

type TransferType string

const (
	// TRANSFER_TYPE_USER: transfer between users (tip, airdrop)
	TRANSFER_TYPE_USER TransferType = "user_transfer"
	// TRANSFER_TYPE_APPLICATION: application transfer request
	TRANSFER_TYPE_APPLICATION TransferType = "application_transfer"
	// TRANSFER_TYPE_VAULT: vault transfer
	TRANSFER_TYPE_VAULT TransferType = "vault_transfer"
)
