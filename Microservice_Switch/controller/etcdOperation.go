package controller

import (
	"Microservice_Switch/Err"
	"Microservice_Switch/Initialize"
	"Microservice_Switch/server"
	"bytes"
	"context"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"time"
)

func PutEtcdKey(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logs.Warn("read r body err:%v", err)
			ResponseErr(w, Err.ErrorResponseReadFail)
			return
		}
		etcdbody := new(EtcdValues)
		err = json.Unmarshal(body, etcdbody)
		if err != nil {
			logs.Warn("json.Unmarshal r.body err %v", err)
			ResponseErr(w, Err.ErrorJsonFailed)
			return
		}

		var strBuf bytes.Buffer
		strBuf.WriteString(etcdbody.Ip)
		strBuf.WriteString(":")
		strBuf.WriteString(etcdbody.Port)
		etcdvalue := strBuf.Bytes()

		_, err = Initialize.EtcdClient.Put(ctx, etcdbody.Key, string(etcdvalue))
		if err != nil {
			logs.Warn("put etcd key:%v value:%v err:%v", etcdbody.Key, etcdvalue, err)
			ResponseErr(w, Err.ErrorPutEtcd)
			return
		}
		etcdvalues := new(server.EtcdValues)
		etcdvalues.Value = string(etcdvalue)
		etcdvalues.Description = etcdbody.Description
		etcdvalues.PriKey = etcdbody.Key

		etcdkey := new(server.EtcdKey)
		etcdkey.Key = etcdbody.Key
		etcdkey.Description = etcdbody.Description
		etcdinsert := server.NewEtcdOperationDb()
		err = etcdinsert.InsertEtcd(etcdkey, etcdvalues)
		if err != nil {
			logs.Warn("insert into etcd db err:%v", err)
			ResponseErr(w, Err.ErrorInsertDb)
			return
		} else {
			resp := "insert etcd and db ok"
			NormalResponse(w, resp, 201)
			return
		}
	} else if r.Method == http.MethodGet {
		etcdselect := server.NewEtcdOperationDb()
		value, err := etcdselect.SelectEtcd()
		if err != nil {
			logs.Warn("select db etcd value err:%v", err)
			ResponseErr(w, Err.ErrorSelectDb)
			return
		}
		data, err := json.Marshal(value)
		if err != nil {
			logs.Warn("json marshal etcd value err", err)
			ResponseErr(w, Err.ErrorJsonFailed)
			return
		}
		NormalResponse(w, string(data), 201)
		return
	} else if r.Method == http.MethodDelete {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logs.Warn("read r body err:%v", err)
			ResponseErr(w, Err.ErrorResponseReadFail)
			return
		}
		etcdkv := new(Etcd)
		err = json.Unmarshal(body, etcdkv)
		if err != nil {
			logs.Warn("json.Unmarshal r.body err %v", err)
			ResponseErr(w, Err.ErrorJsonFailed)
			return
		}
		etcdkey := server.NewEtcdOperationDb()
		valid, err := etcdkey.SelectKey(etcdkv.Key)
		if err != nil && !valid {
			logs.Warn("select key err:%v", err)
			ResponseErr(w, Err.ErrorSelectDb)
			return
		} else {
			etcdinfo, err := loadEtcd(etcdkv.Key)
			if err != nil {
				logs.Warn("etcd not delete key :%v", err)
				ResponseErr(w, Err.ErrGetEtcd)
				return
			}
			if len(etcdinfo) <= 1 {
				_, err := Initialize.EtcdClient.Delete(ctx, etcdkv.Key)
				if err != nil {
					logs.Warn("etcd elete key err:%v", err)
					ResponseErr(w, Err.ErrGetEtcd)
					return
				} else {
					resp := "删除成功"
					NormalResponse(w, resp, 201)
					return
				}
				etcdkey := server.NewEtcdOperationDb()
				err = etcdkey.DeleteEtcd(etcdkv.Key)
				if err != nil {
					logs.Warn("delete db err: %v", err)
					ResponseErr(w, Err.ErrorDeleteDb)
					return
				}
			} else if len(etcdinfo) > 1 {
				for i := 0; i < len(etcdinfo); i++ {
					for _, v := range etcdkv.Values {
						if etcdinfo[i] == v {
							etcdinfo[i] = etcdinfo[len(etcdinfo)-1]
							etcdinfo = etcdinfo[:len(etcdinfo)-1]
						}
					}
				}
				etcdvalue, err := json.Marshal(etcdinfo)
				if err != nil {
					logs.Warn("json.Unmarshal r.body err %v", err)
					ResponseErr(w, Err.ErrorJsonFailed)
					return
				}
				_, err = Initialize.EtcdClient.Put(ctx, etcdkv.Key, string(etcdvalue))
				if err != nil {
					logs.Warn("insert into etcd err:%v", err)
					ResponseErr(w, Err.ErrorPutEtcd)
					return
				}
				etcdkeyValue := server.NewEtcdOperationDb()

				err = etcdkeyValue.PutEtcd(etcdkv.Key, string(etcdvalue))
				if err != nil {
					logs.Warn("update db err: %v", err)
					ResponseErr(w, Err.ErrorUpdateDb)
					return
				}
				resp := "删除成功"
				NormalResponse(w, resp, 201)
				return
			}
		}
	} else if r.Method == http.MethodPut {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logs.Warn("read r body err:%v", err)
			ResponseErr(w, Err.ErrorResponseReadFail)
			return
		}
		etcdkv := new(Etcd)
		err = json.Unmarshal(body, etcdkv)
		if err != nil {
			logs.Warn("json.Unmarshal r.body err %v", err)
			ResponseErr(w, Err.ErrorJsonFailed)
			return
		}
		etcdkey := server.NewEtcdOperationDb()
		valid, err := etcdkey.SelectKey(etcdkv.Key)
		if err != nil && !valid {
			logs.Warn("select key err:%v", err)
			ResponseErr(w, Err.ErrorSelectDb)
			return
		} else {
			etcdvalues, err := json.Marshal(etcdkv.Values)
			if err != nil {
				logs.Warn("json.Unmarshal r.body err %v", err)
				ResponseErr(w, Err.ErrorJsonFailed)
				return
			}
			_, err = Initialize.EtcdClient.Put(ctx, etcdkv.Key, string(etcdvalues))
			if err != nil {
				logs.Warn("insert into etcd err:%v", err)
				ResponseErr(w, Err.ErrorPutEtcd)
				return
			}
			resp := "修改成功"
			NormalResponse(w, resp, 201)
			return
		}
	}

	ResponseErr(w, Err.ErrorRequest)
	return
}

func SelectOne(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logs.Warn("read r body err:%v", err)
			ResponseErr(w, Err.ErrorResponseReadFail)
			return
		}
		var key string
		err = json.Unmarshal(body, &key)
		if err != nil {
			logs.Warn("json.Unmarshal r.body err %v", err)
			ResponseErr(w, Err.ErrorJsonFailed)
			return
		}

		etcdkey := server.NewEtcdOperationDb()
		valid, err := etcdkey.SelectOne(key)
		if err != nil && valid == nil {
			logs.Warn("select key err:%v", err)
			ResponseErr(w, Err.ErrorSelectDb)
			return
		} else {
			resp, err := json.Marshal(&valid)
			if err != nil {
				ResponseErr(w, Err.ErrorJsonFailed)
				return
			}
			NormalResponse(w, string(resp), 201)
			return
		}
	}
	ResponseErr(w, Err.ErrorRequest)
	return
}

func loadEtcd(key string) (etcdinfo []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := Initialize.EtcdClient.Get(ctx, key)
	if err != nil {
		logs.Warn("etcd get key err :%v", err)
		return
	}

	for k, v := range resp.Kvs {
		logs.Debug("k:%v,v:%v", k, v)
		err = json.Unmarshal(v.Value, &etcdinfo)
		if err != nil {
			logs.Warn("json unmarshal etcd value err:%v", err)
			return
		}
	}
	return
}
