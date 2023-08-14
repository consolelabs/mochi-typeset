package typeset

type TransactionProfileSource string

const (
	// TRANSACTION_PROFILE_SOURCE_MOCHI_BALANCE: Mochi's profile balance
	TRANSACTION_PROFILE_SOURCE_MOCHI_BALANCE TransactionProfileSource = "mochi-balance"
	// TRANSACTION_PROFILE_SOURCE_MOCHI_VAULT: Mochi's vault
	TRANSACTION_PROFILE_SOURCE_MOCHI_VAULT TransactionProfileSource = "mochi-vault"
)
