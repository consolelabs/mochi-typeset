package typeset

import (
	"database/sql/driver"
	"encoding/json"
)

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

func (a *Activity) Scan(src any) error {
	if src == nil {
		return nil
	}
	switch t := src.(type) {
	case string:
		return json.Unmarshal([]byte(t), a)
	case []byte:
		return json.Unmarshal(t, a)
	default:
		return nil
	}
}

// Value implements the driver Valuer interface.
func (n Activity) Value() (driver.Value, error) {
	return json.Marshal(n)
}
