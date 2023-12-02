package repositories

import (
	"Restful_Go/models"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	FindArticles() ([]models.Article, error)
	GetArticle(ID int) (models.Article, error)
	CreateArticle(article models.Article) (models.Article, error)
	UpdateArticle(article models.Article) (models.Article, error)
	DeleteArticle(article models.Article, ID int) (models.Article, error)
}

type articleRepository struct {
	db *gorm.DB
}

func RepositoryArticle(db *gorm.DB) *articleRepository {
	return &articleRepository{db}
}

func (r *articleRepository) FindArticles() ([]models.Article, error) {
	var articles []models.Article
	err := r.db.Find(&articles).Error

	return articles, err
}

func (r *articleRepository) GetArticle(ID int) (models.Article, error) {
	var article models.Article
	err := r.db.First(&article, ID).Error

	return article, err
}

func (r *articleRepository) CreateArticle(article models.Article) (models.Article, error) {
	err := r.db.Create(&article).Error

	return article, err
}

func (r *articleRepository) UpdateArticle(article models.Article) (models.Article, error) {
	err := r.db.Save(&article).Error

	return article, err
}

func (r *articleRepository) DeleteArticle(article models.Article, ID int) (models.Article, error) {
	err := r.db.Delete(&article, ID).Error

	return article, err
}
