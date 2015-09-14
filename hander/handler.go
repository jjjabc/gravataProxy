package hander

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func Proxy(w http.ResponseWriter, re *http.Request) {
	defer re.Body.Close()
	client := http.DefaultClient
	proxyRequest := *re
	if proxyRequest.URL.Scheme == "" {
		proxyRequest.URL.Scheme = "http"
	}
	/*	if re.URL.Scheme != "http" || re.URL.Scheme != "https" {
		return
	}*/
	isPathOK := strings.HasPrefix(re.URL.Path, "/avatar/")
	if len(re.URL.Path) == 40 {
		isPathOK = isPathOK && true
	} else if (len(re.URL.Path) == 41) && re.URL.Path[40] == '/' {
		isPathOK = isPathOK && true
	}else{
		isPathOK = false
	}
	if !isPathOK {
		http.NotFound(w, re)
		return
	}
	proxyRequest.URL.Host = "gravatar.com"
	proxyRequest.RequestURI = "" // net/http/request.go L216:It is an error to set this field in an HTTP client request.
	response, err := client.Do(&proxyRequest)
	if err != nil {
		http.Error(w, "client.Do error:"+err.Error(), http.StatusInternalServerError)
		return
	}

	bodyRC := response.Body
	defer bodyRC.Close()
	data, err := ioutil.ReadAll(bodyRC)
	if err != nil {
		http.Error(w, "read response body error:"+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
