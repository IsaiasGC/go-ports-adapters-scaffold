package apperror

import "errors"

var (
	ErrContextCanceled = errors.New("context canceled")
)
