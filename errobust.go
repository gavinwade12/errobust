package errobust

import (
	"fmt"
	"net/http"
)

// Basic error codes.
const (
	Conflict  int = http.StatusConflict
	NotFound      = http.StatusNotFound
	Unhandled     = http.StatusInternalServerError
)

// Error is a custom error.
type Error struct {
	msg  string
	code int
}

func (e Error) Error() string {
	return fmt.Sprintf("error[%d]: %s", e.code, e.msg)
}

// Code returns the error's code.
func (e Error) Code() int {
	return e.code
}

// IsConflict returns true if the error has the conflict code.
func (e Error) IsConflict() bool {
	return e.code == Conflict
}

// IsNotFound returns true if the error has the NotFound code.
func (e Error) IsNotFound() bool {
	return e.code == NotFound
}

// IsUnhandled returns true if the error has the Unhandled code.
func (e Error) IsUnhandled() bool {
	return e.code == Unhandled
}

// New returns a new message with the given message and code.
func New(msg string, code int) Error {
	return Error{msg, code}
}

// TryGetCode returns the code if the error can be converted to Error; otherwise,
// an Error will be returned.
func TryGetCode(err error) (int, error) {
	if e, ok := err.(Error); ok {
		return e.code, nil
	}
	return 0, Error{
		msg:  fmt.Sprintf("could not convert type %T to type Error", err),
		code: Unhandled,
	}
}

// GetCode will try to get the code from the error. If a failure occurs, -1
// will be returned.
func GetCode(err error) int {
	code, e := TryGetCode(err)
	if e != nil {
		return -1
	}
	return code
}

// IsConflict returns true if the error can be converted and has the Conflict code.
func IsConflict(err error) bool {
	e, ok := err.(Error)
	return ok && e.IsConflict()
}

// IsNotFound returns true if the error can be converted and has the NotFound code.
func IsNotFound(err error) bool {
	e, ok := err.(Error)
	return ok && e.IsNotFound()
}

// IsUnhandled returns true if the error can be converted and has the Unhandled code.
func IsUnhandled(err error) bool {
	e, ok := err.(Error)
	return ok && e.IsUnhandled()
}

// Handler can be used to embed error handling functionality.
type Handler struct{}

// TryGetCode returns the code if the error can be converted to Error; otherwise,
// an Error will be returned.
func (h Handler) TryGetCode(err error) (int, error) {
	return TryGetCode(err)
}

// GetCode will try to get the code from the error. If a failure occurs, -1
// will be returned.
func (h Handler) GetCode(err error) int {
	return GetCode(err)
}

// IsConflict returns true if the error can be converted and has the Conflict code.
func (h Handler) IsConflict(err error) bool {
	return IsConflict(err)
}

// IsNotFound returns true if the error can be converted and has the NotFound code.
func (h Handler) IsNotFound(err error) bool {
	return IsNotFound(err)
}

// IsUnhandled returns true if the error can be converted and has the Unhandled code.
func (h Handler) IsUnhandled(err error) bool {
	return IsUnhandled(err)
}
