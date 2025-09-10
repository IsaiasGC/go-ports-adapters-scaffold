package httpserver

func (s *Server) BindRoutes() {
	s.echo.GET("/health", s.dependencies.healtHandler.HealthCheck)
	s.echo.GET("/health/dependencies", s.dependencies.healtHandler.DependenciesHealthCheck)

	s.echo.POST("/resume", s.dependencies.userHandler.CreateUser)
	s.echo.GET("/resume/:resumeid", s.dependencies.userHandler.GetUserByID)
}
