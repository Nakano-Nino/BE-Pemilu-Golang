package repositories

import (
	"Restful_Go/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	Getuser(ID int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("Articles").Find(&users).Error

	return users, err
}

func (r *userRepository) Getuser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Articles").First(&user, ID).Error

	return user, err
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}
