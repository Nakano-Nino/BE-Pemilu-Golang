package partaisdto

type CreatePartaiRequest struct {
	Name           string `json:"name" form:"name" validate:"required"`
	Ketum          string `json:"ketum" form:"ketum" validate:"required"`
	VissionMission string `json:"vissionMission" form:"vissionMission" validate:"required"`
	Address        string `json:"address" form:"address" validate:"required"`
	Image          string `json:"image" form:"image" validate:"required"`
	PaslonID       string `json:"paslonId" form:"paslonId" validate:"required"`
}

type UpdatePartaiRequest struct {
	Name           string `json:"name" form:"name"`
	Ketum          string `json:"ketum" form:"ketum"`
	VissionMission string `json:"vissionMission" form:"vissionMission"`
	Address        string `json:"address" form:"address"`
	Image          string `json:"image" form:"image"`
	PaslonID       string `json:"paslonId" form:"paslonId"`
}

type PartaiResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Ketum          string `json:"ketum"`
	VissionMission string `json:"vissionMission"`
	Address        string `json:"address"`
	Image          string `json:"image"`
	PaslonID       int    `json:"paslonId"`
}
