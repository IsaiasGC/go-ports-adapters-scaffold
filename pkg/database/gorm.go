package database

import (
	"fmt"
	"net/url"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	prodEnv  = "prod"
	localEnv = "local"
	DSN_FMT  = "host=%s port=%d user=%s password=%s dbname=%s connect_timeout=%d sslmode=disable"
)

type DataBase interface {
	Create() (*gorm.DB, error)
}

type database struct {
	env    string
	config *Config
}

func NewDataBase(env string, cfg *Config) DataBase {
	return &database{
		env:    env,
		config: cfg,
	}
}

func (d *database) Create() (*gorm.DB, error) {
	connectionString := d.buildConnectionString()

	db, err := d.initializeDBSession(connectionString)
	if err != nil {
		return nil, err
	}

	d.setConnectionPool(db)

	return db, nil
}

func (d *database) buildConnectionString() string {
	return fmt.Sprintf(
		DSN_FMT,
		d.config.Server,
		d.config.Port,
		d.config.User,
		url.QueryEscape(d.config.Password),
		d.config.Database,
		d.config.ConnectTimeOut,
	)
}

func (d *database) initializeDBSession(connectionString string) (*gorm.DB, error) {
	logMode := logger.Info
	if d.env == prodEnv {
		logMode = logger.Silent
	}

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: false,
		},
		Logger: logger.Default.LogMode(logMode),
	})
	if err != nil {
		return nil, err
	}

	if d.env == localEnv {
		db.Debug()
	}
	return db, nil
}

func (d *database) setConnectionPool(db *gorm.DB) {
	psql, err := db.DB()
	if err != nil {
		panic(err)
	}

	psql.SetMaxIdleConns(d.config.MaxIdleConnections)
	psql.SetMaxOpenConns(d.config.MaxOpenConnections)
	psql.SetConnMaxLifetime(time.Minute * time.Duration(d.config.ConnMaxLifetime))
}
