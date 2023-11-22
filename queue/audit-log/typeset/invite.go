package typeset

import "time"

type AuditLogInviteTracker struct {
	Platform                  string    `json:"platform"`
	InviterProfileID          string    `json:"inviter_profile_id"`
	InviteeProfileID          string    `json:"invitee_profile_id"`
	InviterPlatformIdentifier string    `json:"inviter_platform_identifier"`
	InviteePlatformIdentifier string    `json:"invitee_platform_identifier"`
	GuildID                   string    `json:"guild_id"`
	GuildName                 string    `json:"guild_name"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt                 time.Time `json:"updated_at"`
}
