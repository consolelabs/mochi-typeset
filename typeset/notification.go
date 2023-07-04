package typeset

type NotificationType int

// Enum used to identify the type of notification
const (
	// NOTIFICATION_BINANCE_CONNECT: Notification when user connects to Binance
	NOTIFICATION_BINANCE_CONNECT NotificationType = iota
	// NOTIFICATION_BINANCE_DISCONNECT: later
	NOTIFICATION_INTEGRATION_START NotificationType = 1
	// NOTIFICATION_INTEGRATION_FINISH: later
	NOTIFICATION_INTEGRATION_FINISH NotificationType = 2
	// NOTIFICATION_PAY_WITHDRAW: Notification when user withdraws fund from mochi
	NOTIFICATION_PAY_WITHDRAW NotificationType = 3
	// NOTIFICATION_PAY_DEPOSIT: Notification when user deposits fund to mochi
	NOTIFICATION_PAY_DEPOSIT NotificationType = 4
	// NOTIFICATION_PAY_TRANSFER: Notification when user use swap within mochi
	NOTIFICATION_SWAP NotificationType = 5
	// NOTIFICATION_PAY_TRANSFER: Notification when user using vault
	NOTIFICATION_VAULT NotificationType = 6
	// NOTIFICATION_WALLET_CONNECT: Notification when user connect wallet
	NOTIFICATION_WALLET_CONNECT NotificationType = 7
	// NOTIFICATION_TRIGGER: Notification when user trigger
	NOTIFICATION_TRIGGER NotificationType = 8
	// NOTIFICATION_VAULT_PROPOSAL: Notification when user create proposal
	NOTIFICATION_VAULT_PROPOSAL
	// NOTIFICATION_VAULT_VOTE: Notification when user vote proposal
	NOTIFICATION_VAULT_VOTE
)
