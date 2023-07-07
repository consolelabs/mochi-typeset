package typeset

import "errors"

type AuditLogType int

const (
	AuditLogMessageTypeApi AuditLogType = iota
	AuditLogMessageTypeOnchain
	AuditLogMessageTypeTwitter
)

type AuditLogMessage struct {
	Type   AuditLogType `json:"type"`
	ApiLog *AuditLogApi `json:"api_log,omitempty"`
}

func (m *AuditLogMessage) Validate() error {
	switch m.Type {
	case AuditLogMessageTypeApi:
		return m.ApiLog.Validate()
	case AuditLogMessageTypeOnchain:
		return nil
	case AuditLogMessageTypeTwitter:
		return nil
	default:
		return errors.New("invalid audit log message type")
	}
}
