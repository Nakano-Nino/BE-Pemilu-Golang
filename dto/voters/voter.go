package voters

type VoteRequest struct {
	UserID   int `json:"userId" form:"userId" validate:"required"`
	PaslonID int `json:"paslonId" form:"paslonId" validate:"required"`
}

type VoteResponse struct {
	UserID   int `json:"userId"`
	PaslonID int `json:"paslonId"`
}
