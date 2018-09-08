package Response

import (
	"net/http"
	"ApiGateway/Err"
	"github.com/gin-gonic/gin/json"
	"io"
)

func SendErrorResponse(w http.ResponseWriter,errResponse Err.ErrorResponse){
	w.WriteHeader(errResponse.HttpCode)
	errMessage,_ := json.Marshal(&errResponse.Error)
	io.WriteString(w,string(errMessage))
}


func NormalResponse(w http.ResponseWriter,resp string,code int){
	w.WriteHeader(code)
	io.WriteString(w,resp)
}