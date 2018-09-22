package controller

import (
	"LoginModuleServer/server"
	"LoginModuleServer/subassemblyed"
	"github.com/astaxie/beego/logs"
	"log"
	//"LoginModuleServer/config"
	"LoginModuleServer/init"
	"fmt"
	"strings"
	"time"
)

func (r *U) Login(p User, ret *UserTocken) {
	loginUser := server.NewLoginUser()

	users, err := loginUser.SelectUser(p.Email)
	if err != nil {
		logs.Warn("select user err :%v", err)
		ret.Tocken = ""
		ret.Err = fmt.Sprintf("select user err :%v", err)
		return
	}
	if users.Password != subassemblyed.Md5([]byte(p.Password+init.Salts.Salt)) {
		ret.Tocken = ""
		ret.Err = "账号密码错误"
		return
	}
	if users.Status == -1 {
		ret.Tocken = ""
		ret.Err = "该账号被封禁了"
		return
	}

	err = loginUser.InsertUser(time.Now())
	if err != nil {
		logs.Warn("update login time err", err)
		ret.Tocken = ""
		ret.Err = fmt.Sprintf("select user err :%v", err)
		return
	}
	var tocken strings.Builder
	tocken.WriteString(users.Email)
	tocken.WriteString(subassemblyed.Md5([]byte(p.Password + init.Salts.Salt)))
	ret.Tocken = tocken.String()
	ret.Err = ""
}

func ChkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
