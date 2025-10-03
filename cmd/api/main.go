package main

import (
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/cmd/api/di"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/cmd/api/httpserver"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/pkg/logger"
)

// @title Users API
// @version 0.1.0
// @description This is a sample server for a users API.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	err := di.GetContainer().Invoke(func(s *httpserver.Server) {
		s.WithLogger()
		s.WithErrorHandler()
		s.BindRoutes()
		s.WithSawgger()

		if err := s.Start(); err != nil {
			logger.GetLogger().Fatal("server", "Start", err)
		}
	})
	if err != nil {
		logger.GetLogger().Fatal("dig", "Invoke", err)
	}
}
