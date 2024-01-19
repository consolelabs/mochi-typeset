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
	// TRANSACTION_ACTION_REFUND: Refund action (for failed on-chain txns)
	TRANSACTION_ACTION_REFUND TransactionAction = "refund"
	// TRANSACTION_ACTION_VAULT_TRANSFER: Vault transfer action
	TRANSACTION_ACTION_VAULT_TRANSFER TransactionAction = "vault_transfer"
	// TRANSACTION_ACTION_PAYLINK: Paylink action
	TRANSACTION_ACTION_PAYLINK TransactionAction = "paylink"
	// TRANSACTION_ACTION_PAYME: Payme action
	TRANSACTION_ACTION_PAYME TransactionAction = "payme"
	// TRANSACTION_ACTION_SWEEP: Sweep token action
	TRANSACTION_ACTION_SWEEP TransactionAction = "sweep"
	// TRANSACTION_ACTION_INTERNAL_WITHDRAW: Withdraw from application's balance to owner's
	TRANSACTION_ACTION_INTERNAL_WITHDRAW TransactionAction = "internal_withdraw"
	// TRANSACTION_ACTION_INTERNAL_DEPOSIT: Deposit to application from owner's balance
	TRANSACTION_ACTION_INTERNAL_DEPOSIT TransactionAction = "internal_deposit"
	// TRANSACTION_ACTION_EXTERNAL_TRANSFER: Transfer from an external wallet to a mochi profile
	TRANSACTION_ACTION_EXTERNAL_TRANSFER TransactionAction = "external_transfer"
)
