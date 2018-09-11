package server

import (
	"Microservice_Switch/controller"
	"github.com/astaxie/beego/logs"
)

type ApiOperation struct {
}

func NewApiOperation() *ApiOperation {
	return &ApiOperation{}
}

func (p *ApiOperation) SelectApi() (api *Api, err error) {
	sql := "select method,uri,requestargument,responseargument,argumenttype,required from api"
	err = Db.Select(api, sql)
	if err != nil {
		logs.Warn("select api err:%v", err)
		return
	}
	return
}

func (p *ApiOperation) InsertApi(api *controller.Api) (err error) {
	sql := "insert into api(method,uri,requestargument,responseargument,argumenttype,required,explain) values(?,?,?,?,?,?,?)"
	_, err = Db.Exec(sql, api.Method, api.Uri, api.RequestArgument, api.ResponseArgument, api.ArgumentType, api.Required, api.Explain)
	if err != nil {
		logs.Warn("insert db api  err", err)
		return
	}
	return
}
