package controller

import (
	"encoding/json"
	"net/http"
	"projHugo/hugoproxy/proxy/internal/service"
	"projHugo/hugoproxy/proxy/internal/service/models"
)

type Responder interface {
}

type GeoController struct {
	service service.GeoServicer
}

func NewController(service service.GeoServicer) *GeoController {
	return &GeoController{service: service}
}

func (g *GeoController) SearchHandler(w http.ResponseWriter, r *http.Request) {

	var req models.SearchRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		g.RespondWithError(w, err)
	}

	address, err := g.service.Search(req.Query)
	if err != nil {
		g.RespondWithError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	g.RespondWithSuccess(w, address)

}

func (g *GeoController) GeocodeHandler(w http.ResponseWriter, r *http.Request) {
	var req models.GeocodeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		g.RespondWithError(w, err)
	}

	address, err := g.service.Geocode(req.Lat, req.Lng)
	if err != nil {
		g.RespondWithError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	g.RespondWithSuccess(w, address)

}

func (g *GeoController) RespondWithError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}

func (g *GeoController) RespondWithSuccess(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
