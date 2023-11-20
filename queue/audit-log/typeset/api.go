package typeset

import "time"

type AuditLogApi struct {
	// Method: HTTP method GET, POST, PUT, DELETE
	Method string `json:"method,omitempty"`
	// Uri: URI of the request
	Uri string `json:"uri,omitempty"`
	// RequestBody: request body
	RequestBody []byte `json:"request_body,omitempty"`
	// StatusCode: status code of the response
	StatusCode int `json:"status_code,omitempty"`
	// Latency: process time
	Latency time.Duration `json:"latency,omitempty"`
	// RequestId: request id
	RequestId string `json:"request_id,omitempty"`
	// ResponseBody: response body
	ResponseBody []byte `json:"response_body,omitempty"`
	// Host: host of the request
	Host string
}

func (a *AuditLogApi) Validate() error {
	return nil
}
