package models

import "time"

type Paslon struct {
	ID             int              `json:"id" gorm:"primaryKey:autoIncrement"`
	Name           string           `json:"name" gorm:"type: varchar(255)"`
	OrderNum       int              `json:"orderNum" gorm:"type: integer"`
	VissionMission string           `json:"vissionMission" gorm:"type: varchar(255)"`
	Image          string           `json:"image" gorm:"type: varchar(255)"`
	Partais        []PartaiResponse `json:"partais" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Voters         []Voter          `json:"voters" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt      time.Time        `json:"createdAt"`
	UpdatedAt      time.Time        `json:"updatedAt"`
}

type PaslonVoterResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	OrderNum int    `json:"orderNum"`
}

func (PaslonVoterResponse) TableName() string {
	return "paslons"
}
