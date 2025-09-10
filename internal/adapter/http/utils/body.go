package utils

import (
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/apperror"
	echo "github.com/labstack/echo/v4"
)

func GetBody[T any](ctx echo.Context) (*T, error) {
	var body T
	if err := ctx.Bind(&body); err != nil {
		return nil, apperror.NewError(
			apperror.CodeInvalidParams,
			"invalid request body",
			err,
		)
	}
	return &body, nil
}
