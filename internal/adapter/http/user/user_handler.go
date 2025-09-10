package user

import (
	"net/http"

	httputils "github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/adapter/http/utils"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/interfaces"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/pkg/logger"
	echo "github.com/labstack/echo/v4"
)

type UserHandler struct {
	service interfaces.UserService
	logger  logger.Logger
}

func NewUserHandler(s interfaces.UserService, l logger.Logger) *UserHandler {
	return &UserHandler{
		service: s,
		logger:  l,
	}
}

func (h *UserHandler) CreateUser(ctx echo.Context) error {
	userDto, err := httputils.GetBody[UserDTO](ctx)
	if err != nil {
		return err
	}
	err = h.service.Create(fromUserDTO(userDto))
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusCreated)
}

func (h *UserHandler) GetUserByID(ctx echo.Context) error {
	id := ctx.Param("id")

	user, err := h.service.GetByID(id)
	if err != nil {
		return err
	}

	return ctx.JSON(200, toUserDTO(user))
}
