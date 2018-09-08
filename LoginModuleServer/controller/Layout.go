package controller

import (
	"go_dev/day20/api/defs"
	"time"
)

type UserRegister struct {
	Username   string
	Password1  string
	Password2  string
	Sex        string
	Email      string
	LastLogin  time.Time
	Status     int
	LastIp     string
	CreateTime time.Time
	UpdateTime time.Time
}

type Verifiation struct {
	Code          string
	SendEmailTime time.Time
	Email         string
	Username      string
}

type Signed struct {
	Ver      bool
	username string
}

type User struct {
	Email, Password, Ip string
}

type UserTocken struct {
	Tocken string
	Err    string
}

type U struct {
}
