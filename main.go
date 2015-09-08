package main

import "net/http"

func main() {
	http.HandleFunc("/",Proxyto)
	http.ListenAndServe(":8080", nil)
}
