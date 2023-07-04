package typeset

type NotificationType int

// Enum used to identify the type of notification
const (
	// NOTIFICATION_BINANCE_CONNECT: Notification when user connects to Binance
	NOTIFICATION_BINANCE_CONNECT NotificationType = iota
	// NOTIFICATION_BINANCE_DISCONNECT: later
	NOTIFICATION_INTEGRATION_START
	// NOTIFICATION_INTEGRATION_FINISH: later
	NOTIFICATION_INTEGRATION_FINISH
	// NOTIFICATION_PAY_WITHDRAW: Notification when user withdraws fund from mochi
	NOTIFICATION_PAY_WITHDRAW
	// NOTIFICATION_PAY_DEPOSIT: Notification when user deposits fund to mochi
	NOTIFICATION_PAY_DEPOSIT
	// NOTIFICATION_PAY_TRANSFER: Notification when user use swap within mochi
	NOTIFICATION_SWAP
	// NOTIFICATION_PAY_TRANSFER: Notification when user using vault
	NOTIFICATION_VAULT
	// NOTIFICATION_WALLET_CONNECT: Notification when user connect wallet
	NOTIFICATION_WALLET_CONNECT
	// NOTIFICATION_TRIGGER: Notification when user trigger
	NOTIFICATION_TRIGGER
	// NOTIFICATION_VAULT_PROPOSAL: Notification when user create proposal
	NOTIFICATION_VAULT_PROPOSAL
	// NOTIFICATION_VAULT_VOTE: Notification when user vote proposal
	NOTIFICATION_VAULT_VOTE
	// NOTIFICATION_VAULT_VOTE: Notification when user vote proposal
	NOTIFICATION_EARN
)
