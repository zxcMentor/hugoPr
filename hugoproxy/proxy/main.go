package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	cnt "projHugo/hugoproxy/proxy/internal/controller"
	srv "projHugo/hugoproxy/proxy/internal/service"
	rt "projHugo/hugoproxy/proxy/router"
)

//

func main() {
	cl := &http.Client{}
	geoService := srv.NewGeoServicer(cl)
	geoController := cnt.NewController(geoService)
	r := rt.StartRouter(geoController)

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
