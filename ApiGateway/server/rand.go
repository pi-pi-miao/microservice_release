package server

import (
	"math/rand"
)

type Random struct {
	Rand F
}

type F func(etcdlist []string, key ...string) (addr string, err error)

func RandomBalancing(etcdlist []string, key ...string) (addr string, err error) {
	lens := len(etcdlist)
	index := rand.Intn(lens)
	addr = etcdlist[index]
	return
}
