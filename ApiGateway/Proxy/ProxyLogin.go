package Proxy

import (
	"ApiGateway/Config"
	"ApiGateway/Err"
	"ApiGateway/Response"
	"ApiGateway/initialize"
	"ApiGateway/server"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net"
	"net/http"
	"net/rpc"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		loginIp := ""
		ips := r.Header.Get("X-Forwarded-For")
		if len(ips) > 0 {
			ip := strings.Split(ips, ",")
			loginIp = ip[0]
		} else {
			if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
				loginIp = ip
			}
		}

		resp := fmt.Sprintf("get login ip:%s", loginIp)
		Response.NormalResponse(w, resp, 201)
		return
	} else if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			Response.SendErrorResponse(w, Err.ErrorReadBodyFailed)
			return
		}
		user := new(User)
		tocken := &UserTocken{}

		if err = json.Unmarshal(body, user); err != nil {
			Response.SendErrorResponse(w, Err.ErrorJsonFailed)
			return
		}

		ips := r.Header.Get("X-Forwarded-For")
		if len(ips) > 0 {
			ip := strings.Split(ips, ",")
			user.Ip = ip[0]
		} else {
			if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
				user.Ip = ip
			}
		}
		if user.Email != "" && user.Password != "" {
			addr, err := server.Balance(initialize.EtcdBalance.Values[0], initialize.Etcd.Values)
			if err != nil {
				logs.Error("apigateway balance err:%v", err)
				return
			}

			login, err := rpc.Dial(Config.Conn, addr)
			if err != nil {
				Response.SendErrorResponse(w, Err.ErrorRpcConnFailed)
				return
			}

			err = login.Call("U.Login", user, &tocken)
			if err != nil {
				Response.SendErrorResponse(w, Err.ErrorCall)
				return
			}
			if tocken.Err != "" {
				logs.Debug("err%v", tocken.Err)
				Response.SendErrorResponse(w, Err.ErrorRequestFaild)
				return
			}
		}
		logs.Debug("%s", tocken.Tocken)

		http.SetCookie(w, &http.Cookie{
			Name:  user.Email,
			Value: tocken.Tocken,
		})
		res := "登录成功"
		Response.NormalResponse(w, res, 201)
	} else {
		Response.SendErrorResponse(w, Err.ErrorNotRequest)
		return
	}
}
