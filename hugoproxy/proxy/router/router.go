package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"projHugo/hugoproxy/proxy/internal/controller"
)

func StartRouter(controllers *controller.GeoController) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

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
