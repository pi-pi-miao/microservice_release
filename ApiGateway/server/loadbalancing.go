package server

import "fmt"

func Balance(name string, value []string) (addr string, err error) {
	switch name {
	case "random":
		random := new(Random)
		random.Rand = RandomBalancing
		addr, err = random.Rand(value)
		if err != nil {
			err = fmt.Errorf("random balance value err %s", err)
			return
		}
	}
	return
}
