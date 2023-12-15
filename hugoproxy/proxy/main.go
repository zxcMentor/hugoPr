package main

import (
	"context"
	"database/sql"
	"net/http"
	"proxy/internal/controller"
	"proxy/internal/repository"
	"proxy/internal/service"
	"proxy/internal/service/models"
	"proxy/router"
)

//

func main() {
	us := models.NewUser("OG", 07)
	db := &sql.DB{}
	nDb := repository.NewPostgresUserRepository(db)
	nDb.Create(context.Background(), *us)
	cl := &http.Client{}
	geoService := service.NewGeoServicer(cl)
	geoController := controller.NewController(geoService)
	r := router.StartRouter(geoController)

	http.ListenAndServe(":8080", r)
}
