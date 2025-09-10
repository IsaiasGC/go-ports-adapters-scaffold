package apperror

import "errors"

var (
	ErrInvalidEmail      = errors.New("invalid email")
	ErrNameTooShort      = errors.New("name must be at least 3 characters")
	ErrUserNotFound      = errors.New("user not found")
	ErrInsufficientFunds = errors.New("insufficient funds")
)
