package server

import (
	"errors"
	"net/http"
	"github.com/jjjabc/gravataProxy/hander"
	"strconv"
)
func InitServer(){
	var err error
	http.HandleFunc("/", hander.Proxy)
	go StartServer(nil,uint32(8080))

	err = StartServer(nil,uint32(8443), "TSLkey/CA.cer", "TSLkey/CA.key")
	if err != nil {
		panic(err.Error())
	}
}
func StartServer(handler http.Handler,Port uint32, TSLParameters ...string) error {
	var err error
	var isTSLServer bool
	var ParametersError = errors.New("Bad or invalid parameters")
	if TSLParameters != nil {
		if len(TSLParameters) > 2 {
			return ParametersError
		} else {
			isTSLServer = true
		}
	} else {
		isTSLServer = false
	}
	if Port > 65535 {
		return ParametersError
	}
	if isTSLServer {
		err = http.ListenAndServeTLS(":"+strconv.Itoa(Port), TSLParameters[0], TSLParameters[1], handler)
	} else {
		err = http.ListenAndServe(":"+strconv.Itoa(Port), handler)

	}
	return err
}
