package tokengenerator

import (
	"os"
	"time"

	"auth.com/model"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var key = getEnv("JWT_KEY", "explorer_secret_key")
var issuer = getEnv("JWT_ISSUER", "explorer")
var audience = getEnv("JWT_AUDIENCE", "explorer-front.com")

func GenerateAccessToken(user *model.User) (*AuthenticationResponse, error) {
	claims := jwt.MapClaims{
		"jti":      uuid.New().String(),
		"id":       user.Id,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token, err := CreateToken(claims)
	if err != nil {
		return &AuthenticationResponse{}, err
	}

	return &AuthenticationResponse{
		Id:          user.Id,
		AccessToken: token,
	}, nil
}

func CreateToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
