package typeset

type AuditLogMessage struct {
	Type   string       `json:"type"`
	ApiLog *AuditLogApi `json:"api_log,omitempty"`
}
