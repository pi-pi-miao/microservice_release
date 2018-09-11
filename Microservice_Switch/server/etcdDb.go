package server

import (
	"github.com/astaxie/beego/logs"
)

type EtcdOperationDb struct {
}

func NewEtcdOperationDb() *EtcdOperationDb {
	return &EtcdOperationDb{}
}

func (p *EtcdOperationDb) InsertEtcd(etcdkey *EtcdKey, etcdvalues *EtcdValues) (err error) {
	sql := "insert into etcdvalue(value,description,prikey) values(?,?,?)"
	_, err = Db.Exec(sql, etcdvalues.Value, etcdvalues.Description, etcdvalues.PriKey)
	if err != nil {
		logs.Warn("insert db etcdvalue  err", err)
		return
	}
	sqls := "insert into etcdkey values(?,?)"
	_, err = Db.Exec(sqls, etcdkey.Key, etcdkey.Description)
	defer Db.Close()
	if err != nil {
		logs.Warn("insert db etcdkey err", err)
		return
	}
	return
}

func (p *EtcdOperationDb) SelectEtcd() (ectdValue []*EtcdValues, err error) {
	sql := "select id,value,description,prikey from etcdvalue"
	err = Db.Select(&ectdValue, sql)
	defer Db.Close()
	if err != nil {
		logs.Warn("select etcd value err:%v", err)
		return
	}
	return
}

func (p *EtcdOperationDb) SelectKey(key string) (valid bool, err error) {
	sql := "select key from etcdkey where key=?"
	var etcd []*Etcd
	err = Db.Select(etcd, sql, key)
	if err != nil {
		logs.Warn("db select etcd key err:%v", err)
		return
	}
	if len(etcd) == 0 {
		logs.Warn("etcd is nil")
		valid = false
		return
	}
	valid = true
	return
}

func (p *EtcdOperationDb) DeleteEtcd(key string) (err error) {
	sql := "delete from etcdkey where key=?"
	_, err = Db.Exec(sql, key)
	if err != nil {
		logs.Warn("delete db etcdkey  err", err)
		return
	}
	sqls := "delete from etcdvalue where prikey=?"
	_, err = Db.Exec(sqls, key)
	if err != nil {
		logs.Warn("delete db etcdvalue  err", err)
		return
	}
	return
}

func (p *EtcdOperationDb) PutEtcd(key, value string) (err error) {
	sql := "update etcdvalue set value=? where prikey=?"
	_, err = Db.Exec(sql, key, value)
	if err != nil {
		logs.Warn("update db etcdvalue  err", err)
		return
	}
	return
}

func (p *EtcdOperationDb) SelectOne(key string) ([]string, error) {
	etcd := &Etcd{}
	sql := "select value from etcdvalue where prikey=?"
	err := Db.Select(etcd.Value, sql, key)
	if err != nil {
		logs.Warn("db select etcd key err:%v", err)
		return nil, err
	}
	return etcd.Value, nil
}
