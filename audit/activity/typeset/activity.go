package typeset

import "encoding/json"

type Activity struct {
	Type            ActivityType  `json:"type"`
	UserProfileId   string        `json:"user_profile_id"`
	TargetProfileId string        `json:"target_profile_id"`
	Content         string        `json:"content"`
	Changes         []StateChange `json:"changes"`
	Metadata        interface{}   `json:"metadata"`
}

type StateChange struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (a *Activity) Scan(src interface{}) error {
	if src == nil {
		*a = Activity{}
		return nil
	}
	switch t := src.(type) {
	case string:
		return json.Unmarshal([]byte(t), src)
	case []byte:
		return json.Unmarshal(t, src)
	default:
		return nil
	}
}
