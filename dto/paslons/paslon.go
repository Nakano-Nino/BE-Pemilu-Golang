package paslonsdto

type CreatePaslonRequest struct {
	Name           string `json:"name" form:"name" validate:"required"`
	OrderNum       string `json:"orderNum" form:"orderNum" validate:"required"`
	VissionMission string `json:"vissionMission" form:"vissionMission" validate:"required"`
	Image          string `json:"image" form:"image" validate:"required"`
}

type UpdatePaslonRequest struct {
	Name           string `json:"name" form:"name"`
	OrderNum       string `json:"orderNum" form:"orderNum"`
	VissionMission string `json:"vissionMission" form:"vissionMission"`
	Image          string `json:"image" form:"image"`
}

type PaslonResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	OrderNum       int    `json:"orderNum"`
	VissionMission string `json:"vissionMission"`
	Image          string `json:"image"`
}
