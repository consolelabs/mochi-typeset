package typeset

// ApiPagination: pagination metadata
type ApiPagination struct {
	// Page: the current page
	Page int64 `json:"page"`
	// Size: the number of items per page
	Size int64 `json:"size"`
	// Total: the total number of items
	Total int64 `json:"total"`
}
