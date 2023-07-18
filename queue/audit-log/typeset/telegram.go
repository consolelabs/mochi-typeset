package typeset

type AuditLogTelegram struct {
	Message struct {
		MessageId int64 `json:"message_id,omitempty"`
		From      struct {
			Id       int64  `json:"id,omitempty"`
			IsBot    bool   `json:"is_bot,omitempty"`
			Username string `json:"username,omitempty"`
		} `json:"from,omitempty"`
		Chat struct {
			Id    int64  `json:"id,omitempty"`
			Title string `json:"title,omitempty"`
			Type  string `json:"type,omitempty"`
		} `json:"chat,omitempty"`
		Date int64  `json:"date,omitempty"`
		Text string `json:"text,omitempty"`
	} `json:"message,omitempty"`
}

func (a *AuditLogApi) AuditLogTelegram() error {
	return nil
}
