package typeset

type TransactionAction string

const (
	// TRANSACTION_ACTION_DEPOSIT: Deposit action
	TRANSACTION_ACTION_DEPOSIT TransactionAction = "deposit"
	// TRANSACTION_ACTION_WITHDRAW: Withdraw action
	TRANSACTION_ACTION_WITHDRAW TransactionAction = "withdraw"
	// TRANSACTION_ACTION_TRANSFER: Transfer action
	TRANSACTION_ACTION_TRANSFER TransactionAction = "transfer"
	// TRANSACTION_ACTION_SWAP: Swap action
	TRANSACTION_ACTION_SWAP TransactionAction = "swap"
	// TRANSACTION_ACTION_AIRDROP: Airdrop action
	TRANSACTION_ACTION_AIRDROP TransactionAction = "airdrop"
)
