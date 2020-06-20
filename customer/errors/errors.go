package errors

import (
	"fmt"
)

// Error is a struct
type Error struct {
	Code        int
	Message     string
	OriginError error
}

// Error is a return string error
func (e *Error) Error() string {
	return fmt.Sprintf("Code: %#v Message:%s, Original_error:%s", e.Code, e.Message, e.OriginError.Error())
}
