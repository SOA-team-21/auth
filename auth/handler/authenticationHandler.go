package handler

import (
	"context"

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
