package service

// GeoServicer предоставляет интерфейс для работы с геопространственными запросами.
type GeoServicer interface {
	Search(query string) (SearchResult, error)          // Поиск по заданному запросу.
	SendGeoCoordinates(query string) (GeoResult, error) // Отправка географических координат.
}

// Searcher предоставляет интерфейс для выполнения поисковых запросов.
type Searcher interface {
	Search(query string) (SearchResult, error) // Выполнение поиска и возврат результатов.
}

// CoordinatesServicer предоставляет интерфейс для работы с географическими координатами.
type CoordinatesServicer interface {
	SendGeoCoordinates(query string) (GeoResult, error) // Отправка запроса на получение координат.
}

// Identifier предоставляет интерфейс для аутентификации и генерации токенов.
type Identifier interface {
	GenerateToken(credentials Credentials) (string, error) // Генерация токена на основе предоставленных учетных данных.
}
