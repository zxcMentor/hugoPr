package router

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"net/http/httputil"
	"net/url"
	"proxy/internal/controller"

	"net/http"
)

func StartRouter(controllers *controller.GeoController) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	pr := NewReverseProxy("hugo", "1313")
	r.Use(pr.ReverseProxy)
	r.Get("/swaggers", SwaggerUI)

	r.Post("/swagger", func(writer http.ResponseWriter, request *http.Request) {
		http.StripPrefix("/static", http.FileServer(http.Dir("./"))).ServeHTTP(writer, request)
	})

	rout := chi.NewRouter()
	rout.Post("/address/search", controllers.SearchHandler)
	rout.Post("/address/geocode", controllers.GeocodeHandler)

	r.Mount("/api", rout)
	return r
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
