package hander

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

func TestHanlder(t *testing.T) {
	var Path = ""
	Convey("testing pathCheck", t, func() {
		Convey(`Path为"/"`, func() {
			Path = "/"
			So(pathCheck(Path), ShouldEqual, false)
		})
		Convey(`Path为/test`, func() {
			Path = "/test"
			So(pathCheck(Path), ShouldEqual, false)
		})
		Convey(`Path为/avatar/18ddf52ec2bbc95511fcab6b8a16dd8f3,长于md5hash结果`, func() {
			Path = "/avatar/18ddf52ec2bbc95511fcab6b8a16dd8f3"
			So(pathCheck(Path), ShouldEqual, false)
		})
		Convey(`Path为/avatar/18ddf52ec2bbc95511fcab6b8a16dd8,短于md5hash结果`, func() {
			Path = "/avatar/18ddf52ec2bbc95511fcab6b8a16dd8"
			So(pathCheck(Path), ShouldEqual, false)
		})
		Convey(`Path为/avatar/18ddf52ec2bbc95511fcab6b8a16dd8f/`, func() {
			Path = "/avatar/18ddf52ec2bbc95511fcab6b8a16dd8f/"
			So(pathCheck(Path), ShouldEqual, true)
		})
		Convey(`Path为/avatar/18ddf52ec2bbc95511fcab6b8a16dd8f`, func() {
			Path = "/avatar/18ddf52ec2bbc95511fcab6b8a16dd8f"
			So(pathCheck(Path), ShouldEqual, true)
		})
	})
}

type responseWriter struct {
	*int
}

func (this responseWriter) Header() http.Header {
	return map[string][]string{
		"Accept-Language":        {"en-us"},
		"Connection":             {"keep-alive"},
		"Content-Type":           {"text/plain; charset=utf-8"},
		"X-Content-Type-Options": {"nosniff"},
	}
}
func (this responseWriter) Write(in []byte) (int, error) {
	*(this.int) = len(in)
	return len(in), nil
}
func (this responseWriter) WriteHeader(int) {
	return
}

type readerCloser int

func (this readerCloser) Read(p []byte) (n int, err error) {
	result := "test reader"
	p = []byte(result)
	return len([]byte(result)), nil
}
func (this readerCloser) Close() error {
	return nil
}
func TestProxy(t *testing.T) {
	Convey("测试Proxy", t, func() {
		rw := responseWriter{int: new(int)}
		q, err := http.NewRequest("GET", "http://loaclhost:8080/avatar/18ddf52ec2bbc95511fcab6b8a16dd8f?s=128", new(readerCloser))
		if err != nil {
			Println(err.Error())
			return
		}
		Proxy(rw, q)
		So(*(rw.int), ShouldEqual, 5642)
	})

}
