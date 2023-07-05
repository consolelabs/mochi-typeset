package typeset

import "fmt"

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s", e.Code, e.Message)
}

const (
	CODE_FAIL_VALIDATION = "ERROR_10001"
)

var (
	ERROR_FAIL_VALIDATION = Error{CODE_FAIL_VALIDATION, "Validation Failed"}
)
