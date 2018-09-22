package initialize

import (
	"ApiGateway/Config"
	"ApiGateway/server"
	"context"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	etcd_client "github.com/coreos/etcd/clientv3"
	"github.com/henrylee2cn/tp-micro/discovery/etcd"
	"time"
)

var EtcdClient *etcd_client.Client
var Etcd = new(Config.Etcd)
var EtcdBalance = new(Config.Etcd)

func InitEtcd() error {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{Config.EtcdConf.Addr},
		DialTimeout: time.Duration(Config.EtcdConf.Timeout) * time.Second,
	})
	if err != nil {
		logs.Error("connect etcd failed, err:", err)
		return err
	}

	EtcdClient = cli

	return nil
}

func NewEtcd(key string) *Config.Etcd {
	return &Config.Etcd{
		Key: key,
	}
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

func InitBalance(key string) (err error) {
	logs.Debug("etcd start get balance key")
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
		err = json.Unmarshal(v.Value, EtcdBalance.Values)
		if err != nil {
			logs.Error("json unmarshal etcd value err:%v", err)
			return
		}

	}
	logs.Debug("login service value is ", EtcdBalance.Values)
	return
}

func Initall() (err error) {
	err = InitEtcd()
	if err != nil {
		logs.Warn("init etcd err:%v", err)
		return
	}

	Etcd = NewEtcd(Config.LoginServer)
	err = InitServiceAddr(Etcd.Key)
	if err != nil {
		logs.Error("init service addr err:%v", err)
		return
	}

	etcdBalance := NewEtcd(Config.Balance)
	err = InitBalance(etcdBalance.Key)
	if err != nil {
		logs.Error("init balance err:%v", err)
		return
	}

	return
}
