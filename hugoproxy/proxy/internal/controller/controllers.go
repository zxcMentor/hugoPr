package controller

import (
	"encoding/json"
	"proxy/internal/service"
	"proxy/internal/service/models"

	"net/http"
)

type Responder interface {
	RespondWithError(w http.ResponseWriter, err error)
	RespondWithSuccess(w http.ResponseWriter, data interface{})
}

type GeoController struct {
	service service.GeoServicer
}

func NewController(service service.GeoServicer) *GeoController {
	return &GeoController{service: service}
}

// SearchHandler handles search requests.
// @Summary Search
// @Description This endpoint processes search requests and returns search results.
// @Tags Search
// @Accept  json
// @Produce  json
// @Param query body searchUnm true "Search Query"
// @Success 200 {object} string "Successfully processed search request"
// @Failure 400 {string} string "Bad Request"
// @Router /search [post]
// NewSearchHandler создаёт новый HTTP хендлер для поисковых запросов.
// Возвращает функцию, которая соответствует сигнатуре http.HandlerFunc.
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

// GeocodeHandler handles geocoding requests.
// @Summary Handle geocoding
// @Description This endpoint processes geocoding requests and returns the result.
// @Tags GeoCoding
// @Accept  json
// @Produce  json
// @Param query body geoCodeUnm true "Geocoding Query"
// @Success 200 {string} string "Successfully processed geocoding"
// @Failure 400 {string} string "Bad Request"
// @Router /geocode [post]
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
