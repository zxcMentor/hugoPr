package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"proxy/internal/service/models"

	"net/http"
)

type GeoServicer interface {
	Search(query string) ([]*models.Address, error)
	Geocode(lat, lon float64) ([]*models.Address, error)
}

type geoServicer struct {
	client *http.Client
}

func NewGeoServicer(client *http.Client) GeoServicer {
	return &geoServicer{
		client: client,
	}
}

const (
	DadataURLsugg = "https://suggestions.dadata.ru/suggestions/api/4_1/rs/suggest/address"
	DadataURLgeo  = "https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address"
	APIKey        = "5086f9aa3d01c20cab4d1477df59cb0f1ab75497"
)

func (g *geoServicer) Search(query string) ([]*models.Address, error) {
	req, err := http.NewRequest("POST", DadataURLsugg, bytes.NewBufferString(query))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+APIKey)

	resp, err := g.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("err do req %v", err)
	}
	defer resp.Body.Close()

	var result models.SearchResponse
	err = json.NewDecoder(resp.Body).Decode(&resp)
	if err != nil {
		return nil, fmt.Errorf("err decode js %v", err)
	}
	return result.Addresses, nil

}

func (g *geoServicer) Geocode(lat, lon float64) ([]*models.Address, error) {
	body := map[string]interface{}{
		"lat": lat,
		"lon": lon,
	}
	jsBody, err := json.Marshal(body)

	req, err := http.NewRequest("POST", DadataURLgeo, bytes.NewBuffer(jsBody))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+APIKey)

	resp, err := g.client.Do(req)
	defer resp.Body.Close()

	var result models.GeocodeResponse
	err = json.NewDecoder(resp.Body).Decode(&resp)
	if err != nil {

	}
	return result.Addresses, nil
}
