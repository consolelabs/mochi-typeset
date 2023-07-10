package typeset

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type StateChangeList []StateChange
type Activity struct {
	Type            ActivityType    `json:"type"`
	UserProfileId   string          `json:"user_profile_id"`
	TargetProfileId string          `json:"target_profile_id"`
	Content         string          `json:"content"`
	Changes         StateChangeList `json:"changes"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

type StateChange struct {
	Key string `json:"key"`
	Val string `json:"value"`
}

func (a *StateChangeList) Scan(src any) error {
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
func (n StateChangeList) Value() (driver.Value, error) {
	return json.Marshal(n)
}
