package Proxy

import (
	"ApiGateway/Err"
	"ApiGateway/Response"
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"net"
	"net/http"
	"net/rpc"
	"regexp"
	"strings"
	"sync"
	"time"
)

var now = time.Now()

func init() {
	EmailCache = &sync.Map{}
	DailChan = make(chan bool, 10)
	EmailIn = make(chan bool, 10)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			Response.SendErrorResponse(w, Err.ErrorReadBodyFailed)
			return
		}
		user := getUser()
		err = json.Unmarshal([]byte(body), user)
		if err != nil {
			Response.SendErrorResponse(w, Err.ErrorJsonFailed)
			return
		}
		fmt.Println(now)

		if user.Password1 != "" {
			if len(user.Password1) < 6 {
				Response.SendErrorResponse(w, Err.ErrorPassword)
				return
			} else if user.Password1 != user.Password2 {
				Response.SendErrorResponse(w, Err.ErrorPasswordNotSame)
				return
			}
		}

		if m, _ := regexp.MatchString("^[a-zA-Z]+$", user.Username); !m {
			Response.SendErrorResponse(w, Err.ErrorUserName)
			return
		}
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, user.Email); !m {
			Response.SendErrorResponse(w, Err.ErrorEmail)
			return
		}
		if now != user.CreateTime {
			Response.SendErrorResponse(w, Err.ErrorRegisterTime)
			return
		}

		ips := r.Header.Get("X-Forwarded-For")
		if len(ips) > 0 {
			ip := strings.Split(ips, ",")
			user.LastIp = ip[0]
		} else {
			if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
				user.LastIp = ip
			}
		}

		Register, err := rpc.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			Response.SendErrorResponse(w, Err.ErrorRpcConnFailed)
			return
		}

		var verification = Verification{}
		reply := Register.Go("U.Register", user, &verification, nil)
		if _, ok := <-reply.Done; !ok {
			Response.SendErrorResponse(w, Err.ErrorMethodFailed)
			return
		} else {
			resp := "邮件发送成功"
			Response.NormalResponse(w, resp, 201)
			EmailCache.Store(user.Email, verification)
		}
	} else {
		Response.SendErrorResponse(w, Err.ErrorNotRequest)
		return
	}
}

func RegisterEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			Response.SendErrorResponse(w, Err.ErrorReadBodyFailed)
			return
		}
		verification := new(Verification)
		err = json.Unmarshal(body, verification)
		if err != nil {
			Response.SendErrorResponse(w, Err.ErrorJsonFailed)
			return
		}
		if emailVerific, ok := EmailCache.Load(verification.Email); !ok {
			Response.SendErrorResponse(w, Err.ErrorNotRequest)
			return
		} else {
			if verification.SendEmailTime.Unix()-emailVerific.(Verification).SendEmailTime.Unix() > int64(3*60*time.Second) {
				Response.SendErrorResponse(w, Err.ErrorTimeOut)
				return
			}
			if verification.Code != emailVerific.(Verification).Code {
				Response.SendErrorResponse(w, Err.ErrorRequestFaild)
				return
			}

			Register, err := rpc.Dial("tcp", "127.0.0.1:8080")
			if err != nil {
				DailChan <- false
				log.Printf("Dial err:%s", err)
			}
			suss := new(Signed)
			suss.Ver = true
			suss.username = verification.Username
			var in bool
			reply := Register.Go("U.Register", suss, in, nil)

			resp := "邮箱验证成功"
			Response.NormalResponse(w, resp, 200)

			if _, ok := <-reply.Done; !ok && !in {
				EmailIn <- false
				log.Printf("insert User:%s Db err err:%s", verification.Username, err)
			}

		}
	} else {
		Response.SendErrorResponse(w, Err.ErrorNotRequest)
		return
	}
}

func RegisterDete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		for i := range DailChan {
			if !i {
				//TODO
			} else {
				break
			}
		}
		for k := range EmailIn {
			if !k {
				//TODO
			} else {
				break
			}
		}
	}
}

func getUser() *UserRegister {
	return &UserRegister{
		LastLogin:  now,
		Status:     0,
		UpdateTime: now,
	}
}
