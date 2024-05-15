package repo

import (
	"auth.com/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *UserRepository) GetByUsername(username string) (model.User, error) {
	user := model.User{}
	dbResult := repo.DatabaseConnection.Find(&user, "username = ?", username)
	if dbResult.Error != nil {
		return user, dbResult.Error
	}
	return user, nil
}
