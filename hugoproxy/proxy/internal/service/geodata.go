package service

type Address struct {
	Value             string                 `json:"value"`
	UnrestrictedValue string                 `json:"unrestricted_value"`
	Data              map[string]interface{} `json:"data"`
}

// SearchRequest is a model for search request parameters
// swagger:parameters searchHandler
type SearchRequest struct {
	// Query is the address to search for
	// in: body
	// required: true
	Query string `json:"query"`
}

// SearchResponse is a model for search response data
// swagger:response searchResponse
type SearchResponse struct {
	// in: addresses
	Addresses []*Address `json:"suggestions"`
}

// GeocodeRequest is a model for search request parameters
// swagger:parameters geocodeHandler
type GeocodeRequest struct {
	// in: body
	// required: true
	Lat float64 `json:"lat"`
	Lng float64 `json:"lon"`
}

// GeocodeResponse is a model for search response data
// swagger:response geocodeResponse
type GeocodeResponse struct {
	// in: addresses
	Addresses []*Address `json:"suggestions"`
}

const (
	DadataURLsugg = "https://suggestions.dadata.ru/suggestions/api/4_1/rs/suggest/address"
	DadataURLgeo  = "https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address"
	APIKey        = "5086f9aa3d01c20cab4d1477df59cb0f1ab75497"
)
