package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"proxy/internal/service"
	"proxy/middleware"
	"proxy/utils"
)

// HandleLogin handles user login requests.
// @Summary User Login
// @Description This endpoint authenticates a user and returns a JWT token.
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param credentials body service.Credentials true "User Credentials"
// @Success 200 {object} string "JWT Token"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 415 {string} string "Unsupported Media Type"
// @Failure 500 {string} string "Internal Server Error"
// @Router /login [post]
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	// Extracting and unmarshalling data
	var extraData utils.ExtractDataFromRequest
	var err error
	extraData, err = extraData.Extract(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "incorrect data type", http.StatusUnsupportedMediaType)
		return
	}

	var credentials service.Credentials
	err = extraData.UnmarshalAndProcess(r, &credentials)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	jwtToken, err := service.AuthenticateUser(credentials)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	jsonResponse, err := json.Marshal(jwtToken)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

type JwtResponseBody struct {
	Token string `json:"token"`
}

func SendJwtResponse(w http.ResponseWriter, jwtToken string) {
	response := middleware.JwtResponse{Body: JwtResponseBody{Token: jwtToken}}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
