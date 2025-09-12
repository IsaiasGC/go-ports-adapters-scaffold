package health

import (
	"context"
	"net/http"
	"time"

	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/config"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/interfaces"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/pkg/logger"
	echo "github.com/labstack/echo/v4"
)

type HealthHandler struct {
	healthTimeout time.Duration
	logger        logger.Logger
	service       interfaces.HealthService
}

func NewHealthHandler(c *config.Configuration, l logger.Logger, s interfaces.HealthService) *HealthHandler {
	return &HealthHandler{
		healthTimeout: c.APIConfig.HealthTimeout,
		service:       s,
		logger:        l,
	}
}

func (h *HealthHandler) HealthCheck(ctx echo.Context) error {
	health := h.service.Check()

	return ctx.JSON(http.StatusOK, toHealthDTO(health))
}

func (h *HealthHandler) DependenciesHealthCheck(ctx echo.Context) error {
	childCtx, cancel := context.WithTimeout(ctx.Request().Context(), h.healthTimeout)
	defer cancel()

	health := h.service.CheckDependencies(childCtx)

	return ctx.JSON(http.StatusOK, toHealthDTO(health))
}
