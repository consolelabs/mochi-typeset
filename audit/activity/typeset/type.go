// docs: define type for activity
// NEW TYPE MUST APPEND AT THE END OF THE LIST
package typeset

type ActivityType int

const (
	ACTIVITY_PROFILE_ADD_ONCHAIN_WALLET ActivityType = iota
	ACTIVITY_PROFILE_REMOVE_ONCHAIN_WALLET
	ACTIVITY_PROFILE_ADD_BINANCE
	ACTIVITY_PROFILE_REMOVE_BINANCE
	ACTIVITY_PROFILE_ADD_TELEGRAM
	ACTIVITY_PROFILE_REMOVE_TELEGRAM
	ACTIVITY_PROFILE_ADD_TWITTER
	ACTIVITY_PROFILE_REMOVE_TWITTER
	ACTIVITY_PAY_DEPOSIT
	ACTIVITY_PAY_WITHDRAW
	ACTIVITY_PAY_SEND
	ACTIVITY_PAY_RECEIVE
	ACTIVITY_PAY_SWAP
	ACTIVITY_PAY_EARN
)