package repositories

import (
	"Restful_Go/models"

	"gorm.io/gorm"
)

type PartaiRepository interface {
	FindPartais() ([]models.Partai, error)
	GetPartai(ID int) (models.Partai, error)
	CreatePartai(partai models.Partai) (models.Partai, error)
	UpdatePartai(partai models.Partai) (models.Partai, error)
	DeletePartai(partai models.Partai, ID int) (models.Partai, error)
}

type partaiRepository struct {
	db *gorm.DB
}

func RepositoryPartai(db *gorm.DB) *partaiRepository {
	return &partaiRepository{db}
}

func (r *partaiRepository) FindPartais() ([]models.Partai, error) {
	var partais []models.Partai
	err := r.db.Find(&partais).Error

	return partais, err
}

func (r *partaiRepository) GetPartai(ID int) (models.Partai, error) {
	var partai models.Partai
	err := r.db.First(&partai, ID).Error

	return partai, err
}

func (r *partaiRepository) CreatePartai(partai models.Partai) (models.Partai, error) {
	err := r.db.Create(&partai).Error

	return partai, err
}

func (r *partaiRepository) UpdatePartai(partai models.Partai) (models.Partai, error) {
	err := r.db.Save(&partai).Error

	return partai, err
}

func (r *partaiRepository) DeletePartai(partai models.Partai, ID int) (models.Partai, error) {
	err := r.db.Delete(&partai, ID).Error

	return partai, err
}
