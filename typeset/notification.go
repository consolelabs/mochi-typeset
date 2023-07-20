// docs: define type for notification
// NEW TYPE MUST APPEND AT THE END OF THE LIST
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
	// NOTIFICATION_EARN: Notification when user earn
	NOTIFICATION_EARN
	// NOTIFICATION_WALLET_TRACKING: Notification when user track wallet
	NOTIFICATION_WALLET_TRACKING
	// NOTIFICATION_SOCIAL_CONNECT: Notification when user connect social account
	NOTIFICATION_SOCIAL_CONNECT
	// NOTIFICATION_PAY_TRANSFER_REQUEST: Notification when user receives new transfer request from an application
	NOTIFICATION_PAY_TRANSFER_REQUEST
)
