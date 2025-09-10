package database

type Config struct {
	Server             string
	Database           string
	User               string
	Password           string
	ConnectTimeOut     int
	MaxOpenConnections int
	MaxIdleConnections int
	ConnMaxLifetime    int
	Port               int
}
