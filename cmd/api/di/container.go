package di

import (
	"sync"

	"github.com/IsaiasGC/poc-ports-adapters-scaffold/cmd/api/httpserver"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/adapter/data"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/adapter/http/health"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/adapter/http/user"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/adapter/messaging"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/application"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/config"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/pkg/logger"

	"go.uber.org/dig"
)

// https://blog.drewolson.org/dependency-injection-in-go

var (
	container *dig.Container
	once      sync.Once
)

func GetContainer() *dig.Container {
	once.Do(func() {
		container = buildContainer()
	})
	return container
}

func buildContainer() *dig.Container {
	c := dig.New()
	if err := registerDependencies(c); err != nil {
		logger.GetLogger().Fatal("container", "buildContainer", err)
	}
	return c
}

func registerDependencies(c *dig.Container) error {
	return diGroupErrors(
		c.Provide(logger.GetLogger),
		c.Provide(config.NewConfiguration),
		c.Provide(buildDB),
		c.Provide(data.NewUserRepository),
		c.Provide(messaging.NewMessageProducer),
		c.Provide(application.NewUserService),
		c.Provide(application.NewHealthService),
		c.Provide(user.NewUserHandler),
		c.Provide(health.NewHealthHandler),
		c.Provide(httpserver.NewServerDependencies),
		c.Provide(httpserver.NewServer),
	)
}

func diGroupErrors(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}
