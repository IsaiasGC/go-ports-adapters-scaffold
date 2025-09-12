package httpserver

import (
	"fmt"
	"net/http"

	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/config"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/pkg/logger"
	echo "github.com/labstack/echo/v4"
)

type Server struct {
	echo         *echo.Echo
	config       *config.Configuration
	logger       logger.Logger
	dependencies *ServerDependencies
}

func NewServer(c *config.Configuration, l logger.Logger, d *ServerDependencies) *Server {
	return &Server{
		echo:         echo.New(),
		logger:       l,
		config:       c,
		dependencies: d,
	}
}

func (s *Server) Start() error {
	return s.echo.Start(fmt.Sprintf(":%s", s.config.APIConfig.Port))
}

func (s *Server) NewServerContext(request *http.Request, writer http.ResponseWriter) echo.Context {
	return s.echo.NewContext(request, writer)
}

func (s *Server) ServerHTTP(writer http.ResponseWriter, request *http.Request) {
	s.echo.ServeHTTP(writer, request)
}
