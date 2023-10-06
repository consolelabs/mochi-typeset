package typeset

type AuditLogLogin struct {
	Platform           string `json:"platform,omitempty"`
	ProfileId          string `json:"profile_id,omitempty"`
	PlatformIdentifier string `json:"platform_identifier,omitempty"`
}
