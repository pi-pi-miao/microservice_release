package Proxy

import (
	"io"
	"net/http"
	"strings"
)

type Api struct {
	Method string
	Addr   string
}

type ApiList struct {
	ApiList []Api
}

func SelectApi(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		api := new(Api)
		var string strings.Builder
		string.WriteString(http.MethodPost)
		string.WriteString("\n")
		method := string.String()
		api.Method = method

		var stringaddr strings.Builder
		stringaddr.WriteString("/login")
		stringaddr.WriteString("\n")
		addr := stringaddr.String()
		api.Addr = addr
		io.WriteString(w, api.Addr)
		io.WriteString(w, api.Method)
	} else {
		io.WriteString(w, "no permission")
	}

}
