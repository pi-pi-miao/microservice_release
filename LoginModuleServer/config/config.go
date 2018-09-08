package config

import "sync"

var (
	Conf       Config
	EmailCache *sync.Map
)

type Config struct {
	MysqlConf MysqlConfig
}

type MysqlConfig struct {
	Username string
	Password string
	Port     int
	Database string
	Host     string
}

type Salts struct {
	Salt string
}
