package httpserver

import (
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/adapter/http/health"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/adapter/http/user"
)

type ServerDependencies struct {
	userHandler  *user.UserHandler
	healtHandler *health.HealthHandler
}

func NewServerDependencies(uh *user.UserHandler, hh *health.HealthHandler) *ServerDependencies {
	return &ServerDependencies{
		userHandler:  uh,
		healtHandler: hh,
	}
}
