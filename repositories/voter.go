package repositories

import (
	"Restful_Go/models"

	"gorm.io/gorm"
)

type VoterRepository interface {
	FindVoters() ([]models.Voter, error)
	GetVoters(ID int) (bool, error)
	Vote(voter models.Voter) (models.Voter, error)
}

type voterRepository struct {
	db *gorm.DB
}

func RepositoryVoter(db *gorm.DB) *voterRepository {
	return &voterRepository{db}
}

func (r *voterRepository) FindVoters() ([]models.Voter, error) {
	var voters []models.Voter
	err := r.db.Find(&voters).Error

	return voters, err
}

func (r *voterRepository) GetVoters(ID int) (bool, error) {
	var voters bool
	err := r.db.First(&voters, ID).Error

	return voters, err
}

func (r *voterRepository) Vote(voter models.Voter) (models.Voter, error) {
	err := r.db.Create(&voter).Error

	return voter, err
}
