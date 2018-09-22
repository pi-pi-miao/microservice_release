package init

import (
	"LoginModuleServer/config"
	"LoginModuleServer/server"
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	etcd_client "github.com/coreos/etcd/clientv3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

var Db *sqlx.DB
var EtcdClient *etcd_client.Client
var Salts *config.Salts
var Etcd = new(config.Etcd)
var Mysqld = []*config.MysqlConfig{}

func init() {
	Salts = &config.Salts{}
	Salts.Salt = "b4GdoZ$&2V7SHk4HLQfJM2vpwLQtLfk34U4*NDp42iL%V@ZFR5OVF$Xl2WK$A4zc"
}

func InitDb() error {
	d := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Mysqld[0].Username,
		Mysqld[0].Password, Mysqld[0].Host, Mysqld[0].Port,
		Mysqld[0].Database)
	database, err := sqlx.Open("mysql", d)
	if err != nil {
		fmt.Println("sqlx open mysql err", err)
		return err
	}
	Db = database
	return nil
}

func InitEtcd() error {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{config.Conf.EtcdConf.Addr},
		DialTimeout: time.Duration(config.Conf.EtcdConf.Timeout) * time.Second,
	})
	if err != nil {
		logs.Error("connect etcd failed, err:", err)
		return err
	}

	EtcdClient = cli

	return nil
}

func NewEtcd(key string) *config.Etcd {
	return &config.Etcd{
		Key: key,
	}
}

func InitDbAddr(key string) (err error) {
	logs.Debug("etcd start get db addr")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := EtcdClient.Get(ctx, key)
	if err != nil {
		logs.Error("etcd get db addr err%v", err)
		return
	}

	for k, v := range resp.Kvs {
		logs.Debug("get from etcd db key:%v,v:%v", k, v)
		err = json.Unmarshal(v.Value, Mysqld)
		if err != nil {
			logs.Error("json unmarshal mysql err:%v", err)
			return
		}
	}
	return
}

func InitServiceAddr(key string) (err error) {
	logs.Debug("etcd start get login key")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := EtcdClient.Get(ctx, key)
	if err != nil {
		logs.Error("etcd client get key err")
		return
	}

	logs.Debug("get etcd login key succ")
	for k, v := range resp.Kvs {
		logs.Debug("key[%v],value[%v]", k, v)
		err = json.Unmarshal(v.Value, Etcd.Values)
		if err != nil {
			logs.Error("json unmarshal etcd value err:%v", err)
			return
		}
	}
	logs.Debug("login service value is ", Etcd.Values)
	return
}

func verdict(addr string) (err error) {
	for _, v := range Etcd.Values {
		if v != addr {
			err = fmt.Errorf("this  LoginModuleService  err ,warning warning!!!")
			logs.Error(err)
			panic("this LoginModuleService err")
		}
	}
}

func Initialize() (err error) {
	err = InitEtcd()
	if err != nil {
		logs.Warn("init etcd err:%v", err)
		return
	}

	Etcd = NewEtcd("login")
	err = InitServiceAddr(Etcd.Key)
	if err != nil {
		logs.Error("init service addr err:%v", err)
		return
	}

	mysql := NewEtcd("mysqladdr")
	err = InitDbAddr(mysql.Key)
	if err != nil {
		logs.Error("init mysql addr err:%v", err)
		return
	}

	err = InitDb()
	if err != nil {
		logs.Error("init err:%v", err)
		return
	}

	server.Init(Db)

	verdict(config.ServiceAddr)
	return nil
}
