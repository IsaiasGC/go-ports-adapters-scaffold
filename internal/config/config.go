package config

import "time"

type Configuration struct {
	ProjectInfo *ProjectInfo
	APIConfig   *APIConfig
	DBConfig    *DBConfig
	KafkaConfig *KafkaConfig
}

type ProjectInfo struct {
	Name        string `env:"PROJECT_NAME,default=ports-adapters"`
	Version     string `env:"VERSION,default=0.1.0"`
	Environment string `env:"GO_ENV,required"`
}

type APIConfig struct {
	Port          string        `env:"PORT,default=8080"`
	TimeOut       time.Duration `env:"API_TIMEOUT,default=300ms"`
	HealthTimeout time.Duration `env:"HEALTHCHECK_TIMEOUT,default=5s"`
}

type DBConfig struct {
	Server             string `env:"DB_SERVER,required"`
	Database           string `env:"DB_NAME,required"`
	User               string `env:"DB_USER,required"`
	Password           string `env:"DB_PASSWORD,required"`
	ConnectTimeOut     int    `env:"DB_CONNECT_TIMEOUT,default=45"`
	MaxOpenConnections int    `env:"DB_MAX_OPEN_CONNECTIONS,default=5"`
	MaxIdleConnections int    `env:"DB_MAX_IDLE_CONNECTIONS,default=30"`
	ConnMaxLifetime    int    `env:"DB_CONN_MAX_LIFETIME,default=60"`
	Port               int    `env:"DB_PORT,default=5432"`
}

type KafkaConfig struct {
	Brokers     []string `env:"KAFKA_BROKERS,required"`
	HealthTopic string   `env:"KAFKA_HEALTH_TOPIC,default=healtcheck-topic"`
}
