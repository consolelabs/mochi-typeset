package typeset

// ApiResponse: a generic response for all API calls
type ApiResponse struct {
	// Data: the actual data
	Data interface{} `json:"data"`
	// Metadata: pagination metadata
	Metadata ApiPagination `json:"metadata"`
}
