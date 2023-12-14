package service

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type GeoServicer interface {
	Search(query string) ([]*Address, error)
	Geocode(lat, lon float64) ([]*Address, error)
}

type geoServicer struct {
	client *http.Client
}

func NewGeoServicer(client *http.Client) GeoServicer {
	return &geoServicer{
		client: client,
	}
}

func (g *geoServicer) Search(query string) ([]*Address, error) {
	req, err := http.NewRequest("POST", DadataURLsugg, bytes.NewBufferString(query))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+APIKey)

	resp, err := g.client.Do(req)
	defer resp.Body.Close()

	var result SearchResponse
	err = json.NewDecoder(resp.Body).Decode(&resp)
	if err != nil {

	}
	return result.Addresses, nil

}

func (g *geoServicer) Geocode(lat, lon float64) ([]*Address, error) {
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

	var result GeocodeResponse
	err = json.NewDecoder(resp.Body).Decode(&resp)
	if err != nil {

	}
	return result.Addresses, nil
}
