package health

import (
	"net/http"

	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/interfaces"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/pkg/logger"
	echo "github.com/labstack/echo/v4"
)

type HealthHandler struct {
	service interfaces.HealthService
	logger  logger.Logger
}

func NewHealthHandler(s interfaces.HealthService, l logger.Logger) *HealthHandler {
	return &HealthHandler{
		service: s,
		logger:  l,
	}
}

func (h *HealthHandler) HealthCheck(ctx echo.Context) error {
	health := h.service.Check()

	return ctx.JSON(http.StatusOK, toHealthDTO(health))
}

func (h *HealthHandler) DependenciesHealthCheck(ctx echo.Context) error {
	health := h.service.CheckDependencies()

	return ctx.JSON(http.StatusOK, toHealthDTO(health))
}
