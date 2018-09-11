package controller

type EtcdValues struct {
	Key         string
	Ip          string
	Port        string
	Description string
}

type Etcd struct {
	Key    string
	Values []string
}

type Api struct {
	Method           string
	Uri              string
	RequestArgument  string
	ResponseArgument string
	ArgumentType     string
	Required         int
	Explain          string
}
