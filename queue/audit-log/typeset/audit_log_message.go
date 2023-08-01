package typeset

import "errors"

type AuditLogType int

const (
	AUDIT_LOG_MESSAGE_TYPE_API AuditLogType = iota
	AUDIT_LOG_MESSAGE_TYPE_ONCHAIN
	AUDIT_LOG_MESSAGE_TYPE_TWITTER
	AUDIT_LOG_MESSAGE_TYPE_TELEGRAM
	AUDIT_LOG_MESSAGE_TYPE_DISCORD
)

type AuditLogMessage struct {
	Type        AuditLogType      `json:"type"`
	ApiLog      *AuditLogApi      `json:"api_log,omitempty"`
	TelegramLog *AuditLogTelegram `json:"telegram_log,omitempty"`
	DiscordLog  *AuditLogDiscord  `json:"discord_log,omitempty"`
}

func (m *AuditLogMessage) Validate() error {
	switch m.Type {
	case AUDIT_LOG_MESSAGE_TYPE_API:
		return m.ApiLog.Validate()
	case AUDIT_LOG_MESSAGE_TYPE_ONCHAIN:
		return nil
	case AUDIT_LOG_MESSAGE_TYPE_TWITTER:
		return nil
	case AUDIT_LOG_MESSAGE_TYPE_TELEGRAM:
		return nil
	case AUDIT_LOG_MESSAGE_TYPE_DISCORD:
		return nil
	default:
		return errors.New("invalid audit log message type")
	}
}
