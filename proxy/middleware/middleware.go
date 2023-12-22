package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JwtResponse определяет структуру ответа с JWT токеном.
// swagger:response jwtResponse
type JwtResponse struct {
	// in:body
	Body struct {
		Token string `json:"token"`
	}
}

func JwtCreate() string {

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //Время работы
		Issuer:    "Disp",                                // Кто создал
		Subject:   "USERID",                              // Кто юзер
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("mykey"))
	if err != nil {
		log.Println("jwt create ", err)
		return ""
	}

	return tokenString

}

// ValidateToken - функция проверки валидности токена
func ValidateToken(w http.ResponseWriter, r *http.Request) *jwt.Token {

	/*
		Извлечение Токена из HTTP Заголовка
	*/

	//Получаем строку с Bearer
	authHader := r.Header.Get("Authorization")
	if authHader == "" {
		log.Println("Header empty")
	}

	//Удаляем из строки Bearer
	tokenString := strings.TrimPrefix(authHader, "Bearer ")

	//Передаем токен строкового типа для разбора и проверки на валидность
	token, err := ParserToken(tokenString)
	if err != nil {
		// Если ошибка - отправляем HTTP ответ и возвращаем nil
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		return nil
	}

	return token

}

// ParserToken - разбор токена под структуру jwt.Token
func ParserToken(tokenString string) (*jwt.Token, error) {

	//Отправляем в парсер токен строковго формата и функцию проверки
	// какого типа у нас метод шифрования
	token, err := jwt.Parse(tokenString, validateToken)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// validateToken - проверка приходящего токена в формате структуры jwt.Token
func validateToken(token *jwt.Token) (interface{}, error) {

	//Проверяем, что токен имеет метод шифрования SigningMethodHMAC,
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	//Возвращаем секретный ключ
	return []byte("mykey"), nil
}

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем токен
		if ValidateToken(w, r) == nil {
			return
		}
		next.ServeHTTP(w, r)
	})
}
