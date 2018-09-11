package Initialize

import (
	"Microservice_Switch/server"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/coreos/etcd/clientv3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

var (
	EtcdClient   *clientv3.Client
	ServerConfig = &server.Config{}
	Db           *sqlx.DB
)

func Initialize() (err error) {

	err = initDb()

	err = initEtcd()
	if err != nil {
		logs.Warn("init etcd err%v", err)
		return
	}

	server.Init(Db)

	return
}

func initDb() (err error) {
	d := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", ServerConfig.MysqlConfig.Username,
		ServerConfig.MysqlConfig.Password, ServerConfig.MysqlConfig.Host, ServerConfig.MysqlConfig.Port,
		ServerConfig.MysqlConfig.Database)
	database, err := sqlx.Open("mysql", d)
	if err != nil {
		logs.Warn("sqlx open mysql err", err)
		return err
	}
	Db = database
	return nil
}

func initEtcd() (err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{ServerConfig.EtcdConfig.EtcdAddr},
		DialTimeout: time.Duration(ServerConfig.EtcdConfig.EtcdTimeout) * time.Second,
	})
	if err != nil {
		logs.Error("connect etcd failed, err:", err)
		return
	}

	EtcdClient = cli
	return
}
