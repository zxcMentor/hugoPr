package service

type GeoService struct {
	searchService       Searcher
	coordinatesServicer CoordinatesServicer
}

func NewGeoService(search Searcher, coordinatesService CoordinatesServicer) *GeoService {
	return &GeoService{
		searchService:       search,
		coordinatesServicer: coordinatesService,
	}
}

func (gs *GeoService) Search(query string) (SearchResult, error) {
	return gs.searchService.Search(query)
}

func (gs *GeoService) SendGeoCoordinates(query string) (GeoResult, error) {
	return gs.coordinatesServicer.SendGeoCoordinates(query)
}
