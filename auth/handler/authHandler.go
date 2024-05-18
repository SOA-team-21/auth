package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"auth.com/service"
)

type AuthHandler struct {
	UserService *service.UserService
}

func (handler *AuthHandler) Login(writer http.ResponseWriter, req *http.Request) {
	var credentials service.AuthCredentials
	err := json.NewDecoder(req.Body).Decode(&credentials)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	authResponse, err := handler.UserService.Login(credentials)
	if err != nil {
		println("Error while creating a new Tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(authResponse)
}

func (handler *AuthHandler) ValidateToken(writer http.ResponseWriter, req *http.Request) {
	tokenString := req.Header.Get("Authorization")
	if tokenString == "" {
		println("Error while parsing token")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenString = strings.Split(tokenString, "Bearer")[1]
	tokenString = strings.TrimSpace(tokenString)
	if tokenString == "" {
		println("Cannot retrieve Bearer from token")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err := handler.UserService.ValidateToken(tokenString)
	if err != nil {
		println("Error while validating token")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
