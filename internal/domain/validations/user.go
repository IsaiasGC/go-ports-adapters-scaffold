package validations

import (
	"strings"

	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/apperror"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/models"
)

func ValidateCreateUser(user *models.User) error {
	if !strings.Contains(user.Email, "@") {
		return apperror.NewError(apperror.CodeInvalidParams,
			"validation error",
			apperror.ErrInvalidEmail)
	}
	if len(user.Name) < 3 {
		return apperror.NewError(apperror.CodeInvalidParams,
			"validation error",
			apperror.ErrNameTooShort)
	}
	return nil
}
