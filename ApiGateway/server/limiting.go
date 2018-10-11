package server

import (
	"ApiGateway/Config"
	"github.com/astaxie/beego/logs"
	"sync"
)

var (
	userMap = &sync.Map{}
	ipMap   = &sync.Map{}
	fn      FnSlice
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

type FnSlice []Fn

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
	userMap.Store(id, fn)
	f, _ := userMap.Load(id)
	fnslice := f.(FnSlice)
	fnslice = append(fnslice, userlimit.Sec)
	fnslice = append(fnslice, userlimit.Min)
	userMap.Store(id, fnslice)
}

func IpRegister(ip string) {
	iplimit := new(IpLimit)
	ipMap.Store(ip, fn)
	f, _ := ipMap.Load(ip)
	fnslice := f.(FnSlice)
	fnslice = append(fnslice, iplimit.Sec)
	fnslice = append(fnslice, iplimit.Min)
	ipMap.Store(ip, fnslice)
}

func UserCall(id int, param int64) bool {
	limit, ok := userMap.Load(id)
	if !ok {
		UserRegister(id)
	}
	secUserCount := limit.(FnSlice)[0](param)
	minUserCount := limit.(FnSlice)[1](param)
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
	limit, ok := ipMap.Load(ip)
	if !ok {
		IpRegister(ip)
	}
	secIpCount := limit.(FnSlice)[0](param)
	minIpCount := limit.(FnSlice)[1](param)
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
