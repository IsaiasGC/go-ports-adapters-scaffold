package httpserver

import (
	"os"
	"strings"

	httputils "github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/adapter/http/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func isSkipped(c echo.Context) bool {
	return strings.HasPrefix(c.Path(), "/metrics") ||
		strings.HasPrefix(c.Path(), "/health")
}

func (s *Server) WithLogger() {
	s.echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output:  os.Stdout,
		Skipper: isSkipped,
		Format: `{"time":"${time_rfc3339_nano}", "requestId":"${id}", "method":"${method}", ` +
			`"uri":"${uri}", "status":${status}, "latency_nanoseconds":${latency}, "latency":"${latency_human}"}` +
			"\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))
}

func (s *Server) WithErrorHandler() {
	s.echo.HTTPErrorHandler = func(err error, ctx echo.Context) {
		s.logger.Warningf("httpserver", "HTTPErrorHandler", "handled error: %v", err)

		if e := httputils.HandleHTTPError(ctx, err); e != nil {
			s.logger.Error("httputils", "HandleHTTPError", err)
		}
	}
}
