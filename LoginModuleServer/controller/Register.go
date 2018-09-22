package controller

import (
	"LoginModuleServer/config"
	"LoginModuleServer/init"
	"LoginModuleServer/server"
	"LoginModuleServer/subassemblyed"
	"k8s.io/apimachinery/pkg/util/rand"
	"strconv"
	"sync"
	"time"
)

func init() {
	config.EmailCache = &sync.Map{}
}

func (p *U) Register(r UserRegister, ret *Verifiation) {
	config.EmailCache.Store(r.Username, r)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 4; i++ {
		ret.Code += strconv.Itoa(rand.Intn(10))
	}
	ret.SendEmailTime = subassemblyed.SendEmail(ret.Code, r.Email, r.Username, "发送注册码")
	ret.Email = r.Email
	ret.Username = r.Username
}

func (p *U) UpdateUser(r Signed, ret *bool) {
	if !r.Ver {
		*ret = false
		config.EmailCache.Delete(r.username)
		return
	}
	if UserRegister, ok := config.EmailCache.Load(r.username); !ok {
		*ret = false
		return
	} else {
		var user server.UserRegisterd
		user.Username = UserRegister.(UserRegister).Username
		user.Password1 = subassemblyed.Md5([]byte(UserRegister.(UserRegister).Password1 + init.Salts.Salt))
		user.Password2 = subassemblyed.Md5([]byte(UserRegister.(UserRegister).Password2 + init.Salts.Salt))
		user.Sex = UserRegister.(UserRegister).Sex
		user.Email = UserRegister.(UserRegister).Email
		user.LastLogin = UserRegister.(UserRegister).LastLogin
		user.Status = UserRegister.(UserRegister).Status
		user.LastIp = UserRegister.(UserRegister).LastIp
		user.CreateTime = UserRegister.(UserRegister).CreateTime
		user.UpdateTime = UserRegister.(UserRegister).UpdateTime

		insertUser := server.NewUserRegisterModel()
		err := insertUser.InsertUser(&user)
		if err != nil {
			*ret = false
		}
		*ret = true
	}
}
