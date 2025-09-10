package config

import (
	"context"
	"strings"

	"github.com/IsaiasGC/poc-ports-adapters-scaffold/pkg/envconfig"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/pkg/logger"
	goenvconfig "github.com/sethvargo/go-envconfig"
)

// https://medium.com/@felipedutratine/manage-config-in-golang-to-get-variables-from-file-and-env-variables-33d876887152

func NewConfiguration(log logger.Logger) *Configuration {
	var configuration = Configuration{}
	if err := goenvconfig.Process(context.Background(), &configuration); err != nil {
		log.Fatal("config", "go-envconfig.Process", err)
	}
	if missingEnvVars := envconfig.GetMissingEnvVars(&configuration); len(missingEnvVars) > 0 {
		log.Fatalf(
			"config",
			"envconfig.GetMissingEnvVars",
			"missing required environment variables: %s", strings.Join(missingEnvVars, ", "))
	}
	return &configuration
}
