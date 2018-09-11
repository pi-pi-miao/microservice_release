package controller

import (
	"Microservice_Switch/Err"
	"Microservice_Switch/server"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
)

func SelectApi(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		api := new(Api)
		apioperation := server.NewApiOperation()
		allApi, err := apioperation.SelectApi()
		if err != nil {
			logs.Warn("read r body err:%v", err)
			ResponseErr(w, Err.ErrorSelectDb)
			return
		}
		api.Method = allApi.Method
		api.ArgumentType = allApi.ArgumentType
		api.RequestArgument = allApi.RequestArgument
		api.ResponseArgument = allApi.ResponseArgument
		api.Uri = allApi.Uri
		api.Required = allApi.Required
		api.Explain = allApi.Explain

		resp, err := json.Marshal(api)
		if err != nil {
			ResponseErr(w, Err.ErrorJsonFailed)
			return
		}
		NormalResponse(w, string(resp), 201)
		return
	} else if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logs.Warn("read r body err:%v", err)
			ResponseErr(w, Err.ErrorResponseReadFail)
			return
		}
		api := new(Api)
		err = json.Unmarshal(body, api)
		if err != nil {
			logs.Warn("json.Unmarshal r.body err %v", err)
			ResponseErr(w, Err.ErrorJsonFailed)
			return
		}
		apioperation := server.NewApiOperation()
		err = apioperation.InsertApi(api)
		if err != nil {
			logs.Warn("insert api db err:%v", err)
			ResponseErr(w, Err.ErrorInsertDb)
			return
		}
		resp := "创建api成功"
		NormalResponse(w, string(resp), 201)
		return
	}
	ResponseErr(w, Err.ErrorRequest)
	return
}
