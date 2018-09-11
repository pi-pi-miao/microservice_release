package server

import (
	"net/http"
	"Microservice_Switch/controller"
)

func ServerRouter(){
	http.HandleFunc("/putkey",controller.PutEtcdKey)
	http.HandleFunc("/valid",controller.SelectOne)
	http.HandleFunc("/apioperation",controller.SelectApi)

	http.ListenAndServe(":9000",nil)
}
