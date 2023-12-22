package service

import "errors"

type geoSearch struct {
}

type searchOption func(*geoSearch)

type SearchResult struct {
	Result string
}

func NewSearch(options ...searchOption) *geoSearch {
	var s geoSearch
	for _, option := range options {
		option(&s)
	}
	return &s
}

func (s *geoSearch) Search(query string) (SearchResult, error) {
	if query == "" {
		return SearchResult{}, errors.New("query is empty")
	}
	// Имитация поиска
	return SearchResult{Result: "Все хорошо. Псевдо запрос обработан"}, nil
}
