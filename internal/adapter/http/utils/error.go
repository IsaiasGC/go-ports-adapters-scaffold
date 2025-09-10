package utils

import (
	"errors"
	"net/http"

	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/apperror"
	echo "github.com/labstack/echo/v4"
)

type errorDetail struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

type errorResponse struct {
	Error errorDetail `json:"error"`
}

func HandleHTTPError(ctx echo.Context, err error) error {
	if errors.Is(err, apperror.ErrContextCanceled) {
		return nil
	}

	responseStatus := http.StatusInternalServerError
	httpError := errorDetail{
		Code:    string(apperror.CodeInternalError),
		Message: "internal server error",
	}
	var handledError *apperror.HandledError
	var echoError *echo.HTTPError

	if errors.As(err, &handledError) {
		responseStatus = getStatusFromCodeError(handledError.Code)

		httpError.Code = string(handledError.Code)
		httpError.Message = handledError.Message
		httpError.Description = handledError.Error()
	} else if errors.As(err, &echoError) {
		httpError.Description = echoError.Message.(string)
	} else {
		httpError.Description = err.Error()
	}

	return ctx.JSON(responseStatus, errorResponse{
		Error: httpError,
	})
}

func getStatusFromCodeError(code apperror.CodeError) int {
	switch code {
	case apperror.CodeClientError:
	case apperror.CodeInvalidParams:
		return http.StatusBadRequest
	case apperror.CodeUnauthorized:
		return http.StatusUnauthorized
	case apperror.CodeForbidden:
		return http.StatusForbidden
	case apperror.CodeNotFound:
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
