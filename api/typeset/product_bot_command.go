package typeset

type ProductBotCommandScope int

const (
	PRODUCT_BOT_COMMAND_SCOPE_PUBLIC ProductBotCommandScope = iota
	PRODUCT_BOT_COMMAND_SCOPE_PRIVATE_ONLY
	PRODUCT_BOT_COMMAND_SCOPE_PUBLIC_ONLY
)
