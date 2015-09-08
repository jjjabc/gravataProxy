package hander

import (
	"io/ioutil"
	"net/http"
)

func Proxy(w http.ResponseWriter, re *http.Request) {
	defer re.Body.Close()
	client := http.DefaultClient
	proxyRequest := *re
	if proxyRequest.URL.Scheme == "" {
		proxyRequest.URL.Scheme = "http"
	}
	proxyRequest.URL.Host = "gravatar.com"
	proxyRequest.RequestURI = ""
	response, err := client.Do(&proxyRequest)
	if err != nil {
		w.Write([]byte("client.Do error:" + err.Error()))
		return
	}

	bodyRC := response.Body
	defer bodyRC.Close()
	data, err := ioutil.ReadAll(bodyRC)
	if err != nil {
		w.Write([]byte("read response body error:" + err.Error()))
		return
	}
	w.Write(data)
}
