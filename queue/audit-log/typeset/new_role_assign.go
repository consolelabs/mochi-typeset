package typeset

import "time"

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
