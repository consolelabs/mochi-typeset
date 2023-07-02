package typeset

// Enum used to identify the type of notification
const (
	NotificationBinanceConnect = iota
	NotificationIntegrationStart
	NotificationIntegrationFinish
	NotificationPayWithdraw
	NotificationPayDeposit
	NotificationSwap
	NotificationVault
	NotificationWalletConnect
)
