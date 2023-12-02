package models

type Voter struct {
	ID       int `json:"id" gorm:"primaryKey:autoIncrement"`
	UserID   int `json:"userId"`
	PaslonID int `json:"paslonId"`
}

type VoterUserResponse struct {
	ID       int `json:"id"`
	UserID   int `json:"userId"`
	PaslonID int `json:"paslonId"`
}

type VoterPaslonResponse struct {
	ID       int `json:"id"`
	PaslonID int `json:"paslonId"`
}

type VotersResponse struct {
	Voters      []Voter `json:"voters"`
	VotersCount int     `json:"votersCount"`
}
