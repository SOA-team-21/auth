package service

import (
	"fmt"

	"auth.com/repo"
	tokengenerator "auth.com/tokenGenerator"
)

type UserService struct {
	Repo *repo.UserRepository
}

func (service *UserService) Login(credentials AuthCredentials) (*tokengenerator.AuthenticationResponse, error) {
	authResponse := tokengenerator.AuthenticationResponse{}
	user, err := service.Repo.GetByUsername(credentials.Username)

	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with username %s not found", credentials.Username))
	}
	if !user.IsActive || user.Password != credentials.Password {
		return nil, fmt.Errorf(fmt.Sprintf("credentials invalid or user is not active"))
	}
	authResponse.Id = user.Id
	authResponse.AccessToken = ""
	return &authResponse, nil
}
