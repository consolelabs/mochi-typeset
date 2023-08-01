package typeset

type AuditLogDiscord struct {
	Message struct {
		MessageId string `json:"message_id,omitempty"`
		Author    struct {
			Id       string `json:"id,omitempty"`
			IsBot    bool   `json:"is_bot,omitempty"`
			Username string `json:"username,omitempty"`
		} `json:"from,omitempty"`
		Channel struct {
			Id    string `json:"id,omitempty"`
			Title string `json:"title,omitempty"`
			Type  string `json:"type,omitempty"`
		} `json:"channel,omitempty"`
		Date int64  `json:"date,omitempty"`
		Text string `json:"text,omitempty"`
	} `json:"message,omitempty"`
}

func (a *AuditLogApi) AuditLogDiscord() error {
	return nil
}
