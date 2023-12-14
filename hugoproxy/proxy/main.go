package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"studentgit.kata.academy/zxcMentor/go-kata/course3/2.server/5.server_http_api/task3.2.5.1/hugoproxy/proxy/internal/controller"
	"studentgit.kata.academy/zxcMentor/go-kata/course3/2.server/5.server_http_api/task3.2.5.1/hugoproxy/proxy/internal/service"
	"studentgit.kata.academy/zxcMentor/go-kata/course3/2.server/5.server_http_api/task3.2.5.1/hugoproxy/proxy/router"
)

//

func main() {
	cl := &http.Client{}
	geoService := service.NewGeoServicer(cl)
	geoController := controller.NewController(geoService)
	r := router.StartRouter(geoController)

	http.ListenAndServe(":8080", r)
}

type ReverseProxy struct {
	host string
	port string
}

func NewReverseProxy(host, port string) *ReverseProxy {
	return &ReverseProxy{
		host: host,
		port: port,
	}
}

func (rp *ReverseProxy) ReverseProxy(next http.Handler) http.Handler {
	target, _ := url.Parse(fmt.Sprintf("http://%s:%s", rp.host, rp.port))
	proxy := httputil.NewSingleHostReverseProxy(target)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		proxy.ServeHTTP(w, r)

	})
}
