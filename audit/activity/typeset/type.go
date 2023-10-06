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
	ACTIVITY_APP_WALLET_TRACK
	ACTIVITY_APP_WALLET_UNTRACK
	ACTIVITY_APP_WALLET_ADD
	ACTIVITY_APP_WALLET_ALIAS_SET
	ACTIVITY_APP_WALLET_ALIAS_REMOVE
	ACTIVITY_APP_WATCHLIST_ADD
	ACTIVITY_APP_WATCHLIST_REMOVE
	ACTIVITY_PAY_STAKE
	ACTIVITY_PAY_UNSTAKE
	ACTIVITY_PAY_CLAIM_REWARD
	ACTIVITY_VAULT_PROPOSAL
	ACTIVITY_VAULT_VOTE
	ACTIVITY_PAY_REQUEST
	ACTIVITY_PROFILE_LOGIN
)
