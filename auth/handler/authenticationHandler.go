package handler

import (
	"context"
	"net/http"
	"strings"

	"auth.com/proto/auth"
	"auth.com/service"
)

type AuthenticationHandler struct {
	AuthService *service.UserService
	auth.UnimplementedAuthServiceServer
}

func (handler *AuthenticationHandler) Login(ctx context.Context, request *auth.AuthCredentials) (*auth.AuthenticationResponse, error) {
	username := request.Username
	password := request.Password

	var response, err = handler.AuthService.Login(service.AuthCredentials{Username: username, Password: password})
	if err != nil {
		return &auth.AuthenticationResponse{}, err
	}
	return &auth.AuthenticationResponse{Id: response.Id, AccessToken: response.AccessToken}, nil
}

func (handler *AuthenticationHandler) ValidateToken(ctx context.Context, request *auth.TokenRequest) (*auth.StatusCodeResponse, error) {
	tokenString := request.Token
	if tokenString == "" {
		return &auth.StatusCodeResponse{StatusCode: http.StatusBadRequest}, nil
	}
	tokenString = strings.Split(tokenString, "Bearer")[1]
	tokenString = strings.TrimSpace(tokenString)
	if tokenString == "" {
		return &auth.StatusCodeResponse{StatusCode: http.StatusBadRequest}, nil
	}
	err := handler.AuthService.ValidateToken(tokenString)
	if err != nil {
		return &auth.StatusCodeResponse{StatusCode: http.StatusExpectationFailed}, nil
	}
	return &auth.StatusCodeResponse{StatusCode: http.StatusOK}, nil
}
