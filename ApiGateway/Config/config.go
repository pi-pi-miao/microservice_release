package Config



var(
	Tocken = ""
	Balance = "balance"
	LoginServer = "login"
	Conn = "tcp4"

	Ip_min_acc = 500
	Ip_sec_acc = 50

	User_min_acc =30
	User_sec_acc = 5
)


type EtcdConf struct {
	Addr string
	Timeout int
}

type Etcd struct {
	Key    string
	Values []string
}