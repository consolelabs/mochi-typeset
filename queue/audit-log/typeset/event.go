package typeset

type EventLog struct {
	Type     string      `json:"type"`
	Metadata interface{} `json:"metadata"`
}
