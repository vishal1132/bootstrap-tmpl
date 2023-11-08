package errors

import "fmt"

type Error struct {
	Message    string
	Code       int
	Type       string
	Identifier string
	err        error
}

func (e *Error) Error() string {
	return fmt.Errorf("%w: %s", e.err, e.Message).Error()
}

func NewRecordNotFoundError(err error) *Error {
	return &Error{
		Message:    err.Error(),
		err:        err,
		Code:       404,
		Type:       "DatabaseError",
		Identifier: "RecordNotFound",
	}
}

func NewDatabaseError(err error) *Error {
	return &Error{
		Message: err.Error(),
		err:     err,
		Code:    500,
		Type:    "DatabaseError",
	}
}

func IsRecordNotFoundError(err error) bool {
	v, ok := err.(*Error)
	if !ok {
		return false
	}
	return v.Identifier == "RecordNotFound" && v.Type == "DatabaseError"
}
