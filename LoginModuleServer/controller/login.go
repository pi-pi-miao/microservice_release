package controller

import (
	"LoginModuleServer/config"
	"LoginModuleServer/server"
	"LoginModuleServer/subassemblyed"
	"fmt"
	"github.com/astaxie/beego/logs"
	"log"
	"strings"
	"time"
)

var salts *config.Salts

func init() {
	salts = &config.Salts{}
	salts.Salt = "b4GdoZ$&2V7SHk4HLQfJM2vpwLQtLfk34U4*NDp42iL%V@ZFR5OVF$Xl2WK$A4zc"
}

func (r *U) Login(p User, ret *UserTocken) {
	loginUser := server.NewLoginUser()

	users, err := loginUser.SelectUser(p.Email)
	if err != nil {
		logs.Warn("select user err :%v", err)
		ret.Tocken = ""
		ret.Err = fmt.Sprintf("select user err :%v", err)
		return
	}
	if users.Password != subassemblyed.Md5([]byte(p.Password+salts.Salt)) {
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
	tocken.WriteString(subassemblyed.Md5([]byte(p.Password + salts.Salt)))
	ret.Tocken = tocken.String()
	ret.Err = ""
}

func ChkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
