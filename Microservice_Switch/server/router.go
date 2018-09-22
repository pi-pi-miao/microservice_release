package server

import (
	"Microservice_Switch/controller"
	"net/http"
)

func ServerRouter() {

	http.HandleFunc("/putkey", controller.PutEtcdKey)
	http.HandleFunc("/valid", controller.SelectOne)
	http.HandleFunc("/apioperation", controller.SelectApi)
	http.HandleFunc("/balance", controller.Balance)

	http.ListenAndServe(":9000", nil)
}
