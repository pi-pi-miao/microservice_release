package server

import (
	"ApiGateway/Config"
	"github.com/astaxie/beego/logs"
)

var (
	UserMap = make(map[int][]Fn)
	IpMap   = make(map[string][]Fn)
)

type UserLimit struct {
	count       int
	currentTime int64
}

type IpLimit struct {
	count       int
	currentTime int64
}

type Fn func(nowTime int64) (currentCount int)

func (p *UserLimit) Sec(nowTime int64) (currentCount int) {
	if p.currentTime != nowTime {
		p.count = 1
		p.currentTime = nowTime
		currentCount = p.count
		return
	}

	p.count++
	currentCount = p.count
	return
}

func (p *UserLimit) Min(nowTime int64) (currentCount int) {
	if nowTime-p.currentTime > 60 {
		p.count = 1
		p.currentTime = nowTime
		currentCount = p.count
		return
	}

	p.count++
	currentCount = p.count
	return
}

func (p *IpLimit) Sec(nowTime int64) (currentCount int) {
	if p.currentTime != nowTime {
		p.count = 1
		p.currentTime = nowTime
		currentCount = p.count
		return
	}

	p.count++
	currentCount = p.count
	return
}

func (p *IpLimit) Min(nowTime int64) (currentCount int) {
	if nowTime-p.currentTime > 60 {
		p.count = 1
		p.currentTime = nowTime
		currentCount = p.count
		return
	}

	p.count++
	currentCount = p.count
	return
}

func UserRegister(id int) {
	userlimit := new(UserLimit)
	fn := UserMap[id]
	fn = append(fn, userlimit.Sec)
	fn = append(fn, userlimit.Min)
	UserMap[id] = fn
}

func IpRegister(ip string) {
	iplimit := new(IpLimit)
	fn := IpMap[ip]
	fn = append(fn, iplimit.Sec)
	fn = append(fn, iplimit.Min)
	IpMap[ip] = fn
}

func UserCall(id int, param int64) bool {
	limit, ok := UserMap[id]
	if !ok {
		UserRegister(id)
	}
	secUserCount := limit[0](param)
	minUserCount := limit[1](param)
	if secUserCount > Config.User_sec_acc {
		logs.Warn("this user%v a second visit %v", id, secUserCount)
		return false
	}
	if minUserCount > Config.User_min_acc {
		logs.Warn("this user%v a min visit %v", id, minUserCount)
		return false
	}
	return true
}

func IpCall(ip string, param int64) bool {
	limit, ok := IpMap[ip]
	if !ok {
		IpRegister(ip)
	}
	secIpCount := limit[0](param)
	minIpCount := limit[1](param)
	if secIpCount > Config.Ip_sec_acc {
		logs.Warn("this Ip %v a second visit %v", ip, secIpCount)
		return false
	}
	if minIpCount > Config.Ip_min_acc {
		logs.Warn("this Ip %v a Second visit %v", ip, minIpCount)
		return false
	}
	return true
}
