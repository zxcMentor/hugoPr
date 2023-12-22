package service

import "errors"

// geoCoder представляет собой сервис для преобразования адресов в географические координаты.
type geoCoder struct {
	// Поля, связанные с сервисом геокодирования
}

// GeoResult представляет собой временную структуру для хранения результатов поиска.
type GeoResult struct {
	// как возможный пример формы ответа
	Result string
}

// GeocoderServiceOption определяет тип функции опции для Geocoder.
type GeocoderServiceOption func(*geoCoder)

// NewGeocoder создает новый экземпляр Geocoder с применением переданных опций.
func NewGeocoder(options ...GeocoderServiceOption) *geoCoder {
	var geocoder geoCoder = geoCoder{}

	for _, option := range options {
		option(&geocoder)
	}

	return &geocoder
}

// Geocode выполняет геокодирование на основе предоставленного запроса и возвращает результат.
func (g *geoCoder) Geocode(query string) (GeoResult, error) {
	if query == "" {
		return GeoResult{}, errors.New("query is empty")
	}

	// Имитация процесса геокодирования
	return GeoResult{Result: "Фиктивный результат геокодирования"}, nil
}
