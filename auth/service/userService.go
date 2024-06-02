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
	authResponse := &tokengenerator.AuthenticationResponse{}
	user, err := service.Repo.GetByUsername(credentials.Username)

	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with username %s not found", credentials.Username))
	}
	if !user.IsActive {
		return nil, fmt.Errorf(fmt.Sprintf("user is not active"))
	}

	if user.VerifyPassword(credentials.Password) != nil {
		return nil, fmt.Errorf(fmt.Sprintf("passwords don't match"))
	}

	authResponse, err = tokengenerator.GenerateAccessToken(&user)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("credentials invalid or user is not active"))
	}
	fmt.Printf(fmt.Sprintf("Generated token: %s", authResponse.AccessToken))
	return authResponse, nil
}

func (service *UserService) ValidateToken(tokenString string) error {
	return tokengenerator.ValidateToken(tokenString)
}
