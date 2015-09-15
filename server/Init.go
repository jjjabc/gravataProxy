package server

import (
	"errors"
	"net/http"
	"github.com/jjjabc/gravataProxy/hander"
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
		err = http.ListenAndServeTLS(":8443", TSLParameters[0], TSLParameters[1], handler)
	} else {
		err = http.ListenAndServe(":8080", handler)

	}
	return err
}
