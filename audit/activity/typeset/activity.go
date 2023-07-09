package typeset

type Activity struct {
	Type            ActivityType  `json:"type"`
	UserProfileId   string        `json:"user_profile_id"`
	TargetProfileId string        `json:"target_profile_id"`
	Changes         []StateChange `json:"changes"`
	Metadata        interface{}   `json:"metadata"`
}

type StateChange struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
