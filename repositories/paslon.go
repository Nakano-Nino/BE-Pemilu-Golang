package repositories

import (
	"Restful_Go/models"

	"gorm.io/gorm"
)

type PaslonRepository interface {
	FindPaslons() ([]models.Paslon, error)
	GetPaslon(ID int) (models.Paslon, error)
	CreatePaslon(paslon models.Paslon) (models.Paslon, error)
	UpdatePaslon(paslon models.Paslon) (models.Paslon, error)
	DeletePaslon(paslon models.Paslon, ID int) (models.Paslon, error)
}

type paslonRepository struct {
	db *gorm.DB
}

func RepositoryPaslon(db *gorm.DB) *paslonRepository {
	return &paslonRepository{db}
}

func (r *paslonRepository) FindPaslons() ([]models.Paslon, error) {
	var paslons []models.Paslon
	err := r.db.Preload("Partais").Find(&paslons).Error

	return paslons, err
}

func (r *paslonRepository) GetPaslon(ID int) (models.Paslon, error) {
	var paslon models.Paslon
	err := r.db.Preload("Partais").First(&paslon, ID).Error

	return paslon, err
}

func (r *paslonRepository) CreatePaslon(paslon models.Paslon) (models.Paslon, error) {
	err := r.db.Create(&paslon).Error

	return paslon, err
}

func (r *paslonRepository) UpdatePaslon(paslon models.Paslon) (models.Paslon, error) {
	err := r.db.Save(&paslon).Error

	return paslon, err
}

func (r *paslonRepository) DeletePaslon(paslon models.Paslon, ID int) (models.Paslon, error) {
	err := r.db.Delete(&paslon, ID).Error

	return paslon, err
}
