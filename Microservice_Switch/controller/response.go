package controller

import (
	"Microservice_Switch/Err"
	"encoding/json"
	"io"
	"net/http"
)

func ResponseErr(w http.ResponseWriter, errResponse Err.ErrorResponse) {
	w.WriteHeader(errResponse.HttpCode)
	err_Message, _ := json.Marshal(errResponse.Err)
	io.WriteString(w, string(err_Message))
}

func NormalResponse(w http.ResponseWriter, resp string, httpcode int) {
	w.WriteHeader(httpcode)
	io.WriteString(w, resp)
}
