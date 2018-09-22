package controller

import (
	"Microservice_Switch/Err"
	"Microservice_Switch/Initialize"
	"context"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Balance(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logs.Warn("read r body err:%v", err)
			ResponseErr(w, Err.ErrorSelectDb)
			return
		}
		tactics := new(Etcd)
		err = json.Unmarshal(body, tactics)
		if err != nil {
			logs.Warn("json.Unmarshal r.body err %v", err)
			ResponseErr(w, Err.ErrorJsonFailed)
			return
		}
		var str strings.Builder
		for _, v := range tactics.Values {
			value := v
			str.WriteString(value)
			str.WriteString(",")
		}
		value := str.String()
		_, err = Initialize.EtcdClient.Put(ctx, tactics.Key, value)
		if err != nil {
			logs.Warn("put etcd key:%v value:%v err:%v", tactics.Key, value, err)
			ResponseErr(w, Err.ErrorPutEtcd)
			return
		}

		resp := "insert etcd succ"
		NormalResponse(w, resp, 201)
	} else {
		ResponseErr(w, Err.ErrorRequest)
		return
	}
}
