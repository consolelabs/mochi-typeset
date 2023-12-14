package typeset

import "time"

type AuditLogFeedback struct {
	Platform           string    `json:"platform"`
	ProfileID          string    `json:"profile_id"`
	PlatformIdentifier string    `json:"platform_identifier"`
	Feedback           string    `json:"feedback"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type AuditLogNewRoleAssign struct {
	Platform           string    `json:"platform"`
	ProfileID          string    `json:"profile_id"`
	PlatformIdentifier string    `json:"platform_identifier"`
	GuildID            string    `json:"guild_id"`
	RoleID             string    `json:"role_id"`
	Type               string    `json:"type"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

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

type AuditLogQuestCompleted struct {
	Platform           string    `json:"platform"`
	ProfileID          string    `json:"profile_id"`
	PlatformIdentifier string    `json:"platform_identifier"`
	GuildID            string    `json:"guild_id"`
	QuestID            string    `json:"quest_id"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type AuditLogChannelSetup struct {
	Platform           string    `json:"platform"`
	ProfileID          string    `json:"profile_id"`
	PlatformIdentifier string    `json:"platform_identifier"`
	GuildID            string    `json:"guild_id"`
	ChannelId          string    `json:"channel_id"`
	Action             string    `json:"action"`
	Command            string    `json:"command"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type AuditLogRoleSetup struct {
	Platform           string    `json:"platform"`
	ProfileID          string    `json:"profile_id"`
	PlatformIdentifier string    `json:"platform_identifier"`
	GuildID            string    `json:"guild_id"`
	RoleID             string    `json:"role_id"`
	Action             string    `json:"action"`
	Command            string    `json:"command"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
