package httpserver

import (
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/IsaiasGC/poc-ports-adapters-scaffold/api"
)

func (s *Server) WithSawgger() {
	s.echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
