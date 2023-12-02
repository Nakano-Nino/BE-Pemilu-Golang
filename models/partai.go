package models

import "time"

type Partai struct {
	ID             int       `json:"id" gorm:"primaryKey:autoIncrement"`
	Name           string    `json:"name" gorm:"type: varchar(255)"`
	Ketum          string    `json:"ketum" gorm:"type: varchar(255)"`
	VissionMission string    `json:"vissionMission" gorm:"type: varchar(255)"`
	Address        string    `json:"address" gorm:"type: varchar(255)"`
	Image          string    `json:"image" gorm:"type: varchar(255)"`
	PaslonID       int       `json:"paslonId"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type PartaiResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	PaslonID int    `json:"paslonId"`
}

func (PartaiResponse) TableName() string {
	return "partais"
}
