package main

import (
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/cmd/api/di"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/cmd/api/httpserver"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/pkg/logger"
)

func main() {
	err := di.GetContainer().Invoke(func(s *httpserver.Server) {
		s.WithLogger()
		s.WithErrorHandler()
		s.BindRoutes()

		if err := s.Start(); err != nil {
			logger.GetLogger().Fatal("server", "Start", err)
		}
	})
	if err != nil {
		logger.GetLogger().Fatal("dig", "Invoke", err)
	}
}
