package typeset

type AuditLogDiscord struct {
	Message struct {
		MessageId string `json:"message_id,omitempty"`
		Author    struct {
			Id       string `json:"id,omitempty"`
			IsBot    bool   `json:"is_bot,omitempty"`
			Username string `json:"username,omitempty"`
			IconURL  string `json:"icon_url,omitempty"`
		} `json:"author,omitempty"`
		Channel struct {
			Id    string `json:"id,omitempty"`
			Title string `json:"title,omitempty"`
			Type  string `json:"type,omitempty"`
		} `json:"channel,omitempty"`
		Guild struct {
			Id      string `json:"id,omitempty"`
			Title   string `json:"title,omitempty"`
			Type    string `json:"type,omitempty"`
			IconURL string `json:"icon_url,omitempty"`
		} `json:"guild,omitempty"`
		Date int64  `json:"date,omitempty"`
		Text string `json:"text,omitempty"`
	} `json:"message,omitempty"`
}

func (a *AuditLogApi) AuditLogDiscord() error {
	return nil
}
