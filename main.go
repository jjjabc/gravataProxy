package main

import (
	"github.com/jjjabc/gravataProxy/hander"
	"net/http"
)

func main() {
	http.HandleFunc("/", hander.Proxy)
	http.ListenAndServe(":8080", nil)
}
