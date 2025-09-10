package apperror

import (
	"strings"
)

type CodeError string

const (
	CodeInvalidParams CodeError = "INVALID_PARAMS"
	CodeClientError   CodeError = "CLIENT_ERROR"
	CodeUnauthorized  CodeError = "UNAUTHORIZED"
	CodeForbidden     CodeError = "FORBIDDEN"
	CodeNotFound      CodeError = "NOT_FOUND"
	CodeInternalError CodeError = "INTERNAL_ERROR"
)

type HandledError struct {
	Code    CodeError
	Message string
	errs    []error
}

func NewError(code CodeError, message string, errs ...error) error {
	return &HandledError{
		Code:    code,
		Message: message,
		errs:    errs,
	}
}

func (e *HandledError) Error() string {
	if len(e.errs) == 0 {
		return e.Message
	}

	appErrors := make([]string, len(e.errs))
	for i, err := range e.errs {
		appErrors[i] = err.Error()
	}

	return strings.Join(appErrors, "; ")
}
