package test

import "github.com/jjjabc/gravataProxy/server"

var (
	HTTPSPORT int
	HTTPPORT  int
	HOST      string
)

func init() {
	HTTPSPORT = 8443
	HTTPPORT = 8001
	HOST = "localhost"

	go server.StartServer(nil, uint32(HTTPSPORT), "../TSLkey/CA.cer", "../TSLkey/CA.key")
	go server.StartServer(nil, uint32(HTTPPORT))
}
