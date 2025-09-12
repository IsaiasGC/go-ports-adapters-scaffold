package httpserver

import (
	"fmt"
	"strings"
)

func (s *Server) BindRoutes() {
	s.echo.GET("/health", s.dependencies.healtHandler.HealthCheck)
	s.echo.GET("/health/dependencies", s.dependencies.healtHandler.DependenciesHealthCheck)

	s.echo.POST(toPath(ApiV1, ResourseUsers, ""), s.dependencies.userHandler.CreateUser)
	s.echo.GET(toPath(ApiV1, ResourseUsers, ":id"), s.dependencies.userHandler.GetUserByID)
}

func toPath(api, resource, params string) string {
	path := strings.Join([]string{
		api,
		resource,
		params,
	}, "/")

	path = strings.Trim(path, "/")

	fmt.Println(path)
	return path
}
