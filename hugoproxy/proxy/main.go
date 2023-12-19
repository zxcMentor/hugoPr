package main

import (
	"net/http"
	"proxy/internal/controller"

	"proxy/internal/service"

	"proxy/router"
)

//

func main() {

	cl := &http.Client{}
	geoService := service.NewGeoServicer(cl)
	geoController := controller.NewController(geoService)

	r := router.StartRouter(geoController)

	http.ListenAndServe(":8080", r)
}
