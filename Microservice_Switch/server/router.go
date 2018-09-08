package server

import (
	"Microservice_Switch/controller"
	"net/http"
)

func ServerRouter() {
	http.HandleFunc("", controller.SelectApi)
	http.ListenAndServe(":9000", nil)
}
