package server

import "time"

type UserRegisterd struct {
	Username   string    `db:"username"`
	Password1  string    `db:"password1"`
	Password2  string    `db:"password2"`
	Sex        string    `db:"sex"`
	Email      string    `db:"email"`
	LastLogin  time.Time `db:"lastlogin"`
	Status     int       `db:"status"`
	LastIp     string    `db:"lastip"`
	CreateTime time.Time `db:"createtime"`
	UpdateTime time.Time `db:"updatetime"`
}

type Users struct {
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Lastlogin time.Time `db:"lastlogin"`
	Status    int       `db:"status"`
}
