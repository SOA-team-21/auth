package handler

import (
	"encoding/json"
	"net/http"

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
