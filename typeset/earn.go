package typeset

type Earn struct {
	// Token: Token used for the strategy
	Token Token `json:"token"`
	// Apr: Annual Percentage Rate
	Apr float64 `json:"apr"`
	// Apy: Annual Percentage Yield
	Apy float64 `json:"apy"`
	// Description: Description of the strategy
	Description string `json:"description"`
	// TotalUnderlyingValue: Total underlying assets
	TotalUnderlyingValue TokenAmount `json:"total_underlying_value"`
	// SubscribeInstructions: instrcture how to invest
	SubscribeInstructions *OnchainInstructions `json:"subscribe_instructions"`
	// HarvestInstructions: instrcture how to harvest
	HarvestInstructions *OnchainInstructions `json:"harvest_instructions"`
	// Provider: Provider of the strategy
	Provider string `json:"provider"`
}
