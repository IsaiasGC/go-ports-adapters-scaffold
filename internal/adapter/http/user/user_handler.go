package user

import (
	"context"
	"net/http"
	"time"

	httputils "github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/adapter/http/utils"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/config"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/interfaces"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/pkg/logger"
	echo "github.com/labstack/echo/v4"
)

type UserHandler struct {
	responseTimeout time.Duration
	logger          logger.Logger
	service         interfaces.UserService
}

func NewUserHandler(c *config.Configuration, l logger.Logger, s interfaces.UserService) *UserHandler {
	return &UserHandler{
		responseTimeout: c.APIConfig.TimeOut,
		logger:          l,
		service:         s,
	}
}

func (h *UserHandler) CreateUser(ctx echo.Context) error {
	childCtx, cancel := context.WithTimeout(ctx.Request().Context(), h.responseTimeout)
	defer cancel()

	userDto, err := httputils.GetBody[UserDTO](ctx)
	if err != nil {
		return err
	}

	userModel := fromUserDTO(userDto)
	err = h.service.Create(childCtx, userModel)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, toUserDTO(userModel))
}

func (h *UserHandler) GetUserByID(ctx echo.Context) error {
	childCtx, cancel := context.WithTimeout(ctx.Request().Context(), h.responseTimeout)
	defer cancel()

	id := ctx.Param("id")

	userModel, err := h.service.GetByID(childCtx, id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, toUserDTO(userModel))
}
