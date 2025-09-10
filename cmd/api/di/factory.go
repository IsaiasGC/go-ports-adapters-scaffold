package di

import (
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/config"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/pkg/database"
	"gorm.io/gorm"
)

func buildDB(cfg *config.Configuration) (*gorm.DB, error) {
	dbConfig := &database.Config{
		Server:             cfg.DBConfig.Server,
		Port:               cfg.DBConfig.Port,
		User:               cfg.DBConfig.User,
		Password:           cfg.DBConfig.Password,
		Database:           cfg.DBConfig.Database,
		ConnectTimeOut:     cfg.DBConfig.ConnectTimeOut,
		MaxIdleConnections: cfg.DBConfig.MaxIdleConnections,
		MaxOpenConnections: cfg.DBConfig.MaxOpenConnections,
		ConnMaxLifetime:    cfg.DBConfig.ConnMaxLifetime,
	}
	dbGorm := database.NewDataBase(cfg.ProjectInfo.Environment, dbConfig)

	return dbGorm.Create()
}
