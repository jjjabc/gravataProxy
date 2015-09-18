package test

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"strconv"
	"testing"
	"crypto/tls"
)

func TestHttps(t *testing.T) {
	var (
		Scheme string
		Path   string
		//Query     string
		URLString string
	)
	client := http.DefaultClient
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	var response *http.Response
	var err error
	Convey("测试https", t, func() {
		Scheme = "https"
		URLString = Scheme + "://" + HOST + ":" + strconv.Itoa(HTTPSPORT)
		Convey("测试正常头像链接", func() {
			Convey("Path为/avatar/18ddf52ec2bbc95511fcab6b8a16dd8f", func() {
				Path = "/avatar/18ddf52ec2bbc95511fcab6b8a16dd8f"
				url := URLString + Path
				response, err = client.Get(url)
				So(err, ShouldBeNil)
				So(response.StatusCode, ShouldBeIn, 200, 304)
				So(response.Header.Get("Content-Type"), ShouldEqual, "image/jpeg")
			})
			Convey(`Path为/avatar/18ddf52ec2bbc95511fcab6b8a16dd8f/`, func() {
				Path = "/avatar/18ddf52ec2bbc95511fcab6b8a16dd8f/"
				url := URLString + Path
				response, err = client.Get(url)
				So(err, ShouldBeNil)
				So(response.StatusCode, ShouldBeIn, 200, 304)
				So(response.Header.Get("Content-Type"), ShouldEqual, "image/jpeg")
			})
		})

		Convey("测试非头像链接", func() {
			Convey(`Path为"/"`, func() {
				Path = "/"
				url := URLString + Path
				response, err = client.Get(url)
				So(err, ShouldBeNil)
				So(response.StatusCode, ShouldEqual, 404)
			})
			Convey(`Path为/test`, func() {
				Path = "/test"
				url := URLString + Path
				response, err = client.Get(url)
				So(err, ShouldBeNil)
				So(response.StatusCode, ShouldEqual, 404)
			})
			Convey(`Path为/avatar/18ddf52ec2bbc95511fcab6b8a16dd8f3,长于md5hash结果`, func() {
				Path = "/avatar/18ddf52ec2bbc95511fcab6b8a16dd8f3"
				url := URLString + Path
				response, err = client.Get(url)
				So(err, ShouldBeNil)
				So(response.StatusCode, ShouldEqual, 404)
			})
			Convey(`Path为/avatar/18ddf52ec2bbc95511fcab6b8a16dd8,短于md5hash结果`, func() {
				Path = "/avatar/18ddf52ec2bbc95511fcab6b8a16dd8"
				url := URLString + Path
				response, err = client.Get(url)
				So(err, ShouldBeNil)
				So(response.StatusCode, ShouldEqual, 404)
			})
		})
	})
}
