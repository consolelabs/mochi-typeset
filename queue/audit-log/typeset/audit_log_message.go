package typeset

import "errors"

type AuditLogType int

const (
	AUDIT_LOG_MESSAGE_TYPE_API AuditLogType = iota
	AUDIT_LOG_MESSAGE_TYPE_ONCHAIN
	AUDIT_LOG_MESSAGE_TYPE_TWITTER
	AUDIT_LOG_MESSAGE_TYPE_TELEGRAM
	AUDIT_LOG_MESSAGE_TYPE_DISCORD
	AUDIT_LOG_MESSAGE_TYPE_LOGIN
	AUDIT_LOG_MESSAGE_TYPE_BOT_GM
	AUDIT_LOG_MESSAGE_TYPE_BOT_INVITE_TRACKER
	AUDIT_LOG_MESSAGE_TYPE_BOT_NEW_ROLE_ASSIGN
	AUDIT_LOG_MESSAGE_TYPE_BOT_FEEDBACK
	AUDIT_LOG_MESSAGE_TYPE_BOT_QUEST_COMPLETED
	AUDIT_LOG_MESSAGE_TYPE_BOT_CHANNEL_SETUP
	AUDIT_LOG_MESSAGE_TYPE_BOT_ROLE_SETUP
)

type AuditLogMessage struct {
	Type              AuditLogType            `json:"type"`
	ApiLog            *AuditLogApi            `json:"api_log,omitempty"`
	TelegramLog       *AuditLogTelegram       `json:"telegram_log,omitempty"`
	DiscordLog        *AuditLogDiscord        `json:"discord_log,omitempty"`
	LoginLog          *AuditLogLogin          `json:"login_log,omitempty"`
	GmLog             *AuditLogGM             `json:"gm_log,omitempty"`
	InviteTrackerLog  *AuditLogInviteTracker  `json:"invite_tracker_log,omitempty"`
	NewRoleAssignLog  *AuditLogNewRoleAssign  `json:"new_role_assign_log,omitempty"`
	FeedbackLog       *AuditLogFeedback       `json:"feedback_log,omitempty"`
	QuestCompletedLog *AuditLogQuestCompleted `json:"quest_completed_log,omitempty"`
	ChannelSetupLog   *AuditLogChannelSetup   `json:"channel_setup_log,omitempty"`
	RoleSetupLog      *AuditLogRoleSetup      `json:"role_setup_log,omitempty"`
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
	case AUDIT_LOG_MESSAGE_TYPE_LOGIN:
		return nil
	case AUDIT_LOG_MESSAGE_TYPE_BOT_GM:
		return nil
	case AUDIT_LOG_MESSAGE_TYPE_BOT_INVITE_TRACKER:
		return nil
	case AUDIT_LOG_MESSAGE_TYPE_BOT_NEW_ROLE_ASSIGN:
		return nil
	case AUDIT_LOG_MESSAGE_TYPE_BOT_FEEDBACK:
		return nil
	case AUDIT_LOG_MESSAGE_TYPE_BOT_QUEST_COMPLETED:
		return nil
	case AUDIT_LOG_MESSAGE_TYPE_BOT_CHANNEL_SETUP:
		return nil
	case AUDIT_LOG_MESSAGE_TYPE_BOT_ROLE_SETUP:
		return nil

	default:
		return errors.New("invalid audit log message type")
	}
}
