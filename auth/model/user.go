package model

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role int

const (
	Administrator Role = iota
	Author
	Tourist
)

type User struct {
	Id                 int64 `json:"id"`
	Username           string
	Password           string
	Role               Role
	Email              string
	IsActive           bool
	IsProfileActivated bool
	IsBlogEnabled      bool
}

func (user *User) BeforeCreate(scope *gorm.DB) error {
	err := user.Validate()
	if err != nil {
		return err
	}
	err = user.HashPassword()
	if err != nil {
		return err
	}
	return nil
}

func (user *User) Validate() error {
	if user.Username == "" {
		return errors.New("invalid username")
	}
	if user.Email == "" {
		return errors.New("invalid email")
	}
	if user.Password == "" {
		return errors.New("invalid password")
	}
	return nil
}

func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (r Role) String() string {
	switch r {
	case 0:
		return "administrator"
	case 1:
		return "author"
	case 2:
		return "tourist"
	default:
		return "unknown"
	}
}
