package models

import "time"

type Article struct {
	ID          int       `json:"id" gorm:"primaryKey:autoIncrement"`
	ArticleName string    `json:"articleName" gorm:"type: varchar(255)"`
	Description string    `json:"description" gorm:"type: varchar(255)"`
	Image       string    `json:"image" gorm:"type: varchar(255)"`
	UserID      int       `json:"userId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ArticleResponse struct {
	ID          int    `json:"id"`
	ArticleName string `json:"articleName"`
	Description string `json:"description"`
	Image       string `json:"image"`
	UserID      int    `json:"userId"`
}

func (ArticleResponse) TableName() string {
	return "articles"
}
