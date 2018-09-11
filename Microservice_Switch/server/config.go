package server

type Config struct {
	EtcdConfig  EtcdConfig
	MysqlConfig MysqlConfig
}

type EtcdConfig struct {
	EtcdAddr    string
	EtcdTimeout int
}

type MysqlConfig struct {
	Username string
	Password string
	Port     int
	Database string
	Host     string
}

type EtcdKey struct {
	Key         string `Db:"key"`
	Description string `Db:"description"`
}

type EtcdValues struct {
	Value       string `Db:"value"`
	Description string `Db:"description"`
	PriKey      string `Db:"prikey"`
}

type Etcd struct {
	Key   string   `Db:"key"`
	Value []string `Db:"value"`
}

type Api struct {
	Method           string `Db:"method"`
	Uri              string `Db:"uri"`
	RequestArgument  string `Db:"requestargument"`
	ResponseArgument string `Db:"responseargument"`
	ArgumentType     string `Db:"argumenttype"`
	Required         int    `Db:"required"`
	Explain          string `Db:"explain"`
}
