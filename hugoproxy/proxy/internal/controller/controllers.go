package controller

import (
	"encoding/json"
	"net/http"
	"projHugo/internal/service"
	"projHugo/internal/service/models"
)

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
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	address, err := g.service.Search(req.Query)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(address)

}

func (g *GeoController) GeocodeHandler(w http.ResponseWriter, r *http.Request) {
	var req models.GeocodeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	address, err := g.service.Geocode(req.Lat, req.Lng)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(address)

}
