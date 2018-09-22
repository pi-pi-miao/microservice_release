package Config



var(
	Tocken = ""
	Balance = "balance"
	LoginServer = "login"
	Conn = "tcp4"
)


type EtcdConf struct {
	Addr string
	Timeout int
}

type Etcd struct {
	Key    string
	Values []string
}