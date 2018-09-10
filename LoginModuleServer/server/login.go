package server

import (
	"github.com/labstack/gommon/log"
	"time"
)

type LoginUser struct {
}

func NewLoginUser() *LoginUser {
	return &LoginUser{}
}

func (p *LoginUser) SelectUser(email string) (users *Users, err error) {
	sql := "select password1,email,lastlogin,status from user where email=?"
	err = Db.Select(users, sql, email)

	defer Db.Close()
	if err != nil {
		log.Warn("Select user failed err:%v", err)
		return
	}
	return
}

func (p *LoginUser) InsertUser(lastlogin time.Time) error {
	sql := "insert into user(lastlogin) values(?)"
	_, err := Db.Exec(sql, lastlogin)
	defer Db.Close()
	if err != nil {
		log.Warn("update user lastlogin err", err)
		return err
	}
	return nil
}
