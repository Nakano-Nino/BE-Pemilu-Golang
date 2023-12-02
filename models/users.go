package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey:autoIncrement"`
	Name      string    `json:"name" gorm:"type: varchar(255)"`
	Address   string    `json:"address" gorm:"type: varchar(255)"`
	Gender    string    `json:"gender" gorm:"type: varchar(255)"`
	Username  string    `json:"username" gorm:"type: varchar(255)"`
	Password  string    `json:"password" gorm:"type: varchar(255)"`
	Role      string    `json:"role" gorm:"type: varchar(255)"`
	Articles  []Article `json:"articles"`
	Voter     Voter     `json:"voter"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserVoterResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (UserVoterResponse) TableName() string {
	return "users"
}
