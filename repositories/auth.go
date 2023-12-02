package repositories

import (
	"Restful_Go/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	Login(username string) (models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func RepositoryAuth(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (r *authRepository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *authRepository) Login(username string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "username=?", username).Error

	return user, err
}
