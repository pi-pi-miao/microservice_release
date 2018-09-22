package config

import "sync"

const (
	ServiceAddr = "127.0.0.1:8080"
	Conn        = "tcp4"
)

var (
	Conf       Config
	EmailCache *sync.Map
)

type Config struct {
	MysqlConf MysqlConfig
	EtcdConf  EtcdConf
}

type Etcd struct {
	Key    string
	Values []string
}

type MysqlConfig struct {
	Username string
	Password string
	Port     int
	Database string
	Host     string
}

type EtcdConf struct {
	Addr    string
	Timeout int
}

type Salts struct {
	Salt string
}
