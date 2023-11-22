package typeset

import "time"

type AuditLogGM struct {
	Platform           string    `json:"platform"`
	ProfileID          string    `json:"profile_id"`
	PlatformIdentifier string    `json:"platform_identifier"`
	GuildID            string    `json:"guild_id"`
	ChannelID          string    `json:"channel_id"`
	StreakCount        int       `json:"streak_count"`
	TotalCount         int       `json:"total_count"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
