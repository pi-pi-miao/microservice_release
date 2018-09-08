package server

import (
	"log"
)

type UserRegisterModel struct {
}

func NewUserRegisterModel() *UserRegisterModel {
	return &UserRegisterModel{}
}

func (p *UserRegisterModel) InsertUser(userRegister *UserRegisterd) (err error) {
	sql := "insert into user(username,password1,password2,sex,email,lastlogin,status,lastip,createtime,updatetime)values(?,?,?,?,?,?,?,?,?,?)"
	_, err = Db.Exec(sql, userRegister.Username, userRegister.Password1, userRegister.Password2, userRegister.Sex, userRegister.Email, userRegister.LastLogin, userRegister.Status, userRegister.LastIp, userRegister.CreateTime, userRegister.UpdateTime)
	defer Db.Close()
	if err != nil {
		log.Println("db insert user err", err)
		return
	}
}
