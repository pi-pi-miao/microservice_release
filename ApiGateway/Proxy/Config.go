package Proxy

import (
	"sync"
	"time"
)

var EmailCache *sync.Map

var DailChan chan bool

var EmailIn chan bool

type User struct {
	Email, Password, Ip string
}

type UserTocken struct {
	Tocken string
	Err    string
}

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

type Verification struct {
	Code          string
	SendEmailTime time.Time
	Email         string
	Username      string
}

type Signed struct {
	Ver      bool
	username string
}
