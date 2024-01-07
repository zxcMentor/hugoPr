package router

import (
	"github.com/go-chi/chi/middleware"
	"net/http"
	"proxy/internal/controller"
	"proxy/middlew"

	"github.com/go-chi/chi"
)

type RouterConfig struct {
	router chi.Mux
}

func (rc *RouterConfig) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rc.router.ServeHTTP(w, r)
}

type RouterOption func(*RouterConfig)

func NewRouter(options ...RouterOption) *RouterConfig {

	var router RouterConfig = RouterConfig{}

	for _, option := range options {
		option(&router)
	}

	return &router
}

func PublicRouterOption() RouterOption {
	return func(rc *RouterConfig) {
		var publicRouter *chi.Mux = chi.NewRouter()
		publicRouter.Post("/login", controller.HandleLogin)
		publicRouter.Post("/registration", nil)
		rc.router.Mount("/public", publicRouter)

	}
}

func PrivateRouterOption() RouterOption {
	return func(rc *RouterConfig) {
		var protectedRouter *chi.Mux = chi.NewRouter()
		protectedRouter.Use(middlew.JWTAuthMiddleware)
		protectedRouter.Post("/address/geocode", controller.HandleGeoCode)
		protectedRouter.Post("/address/search", controller.SearchHandler)

		rc.router.Mount("/private", protectedRouter)
	}
}
func SetupRouter() http.Handler {
	router := NewRouter(PublicRouterOption(), PrivateRouterOption())
	return router
}
func StartRout() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/address/search", controller.SearchHandler)
	r.Post("/address/geocode", controller.HandleGeoCode)
	return r
}
