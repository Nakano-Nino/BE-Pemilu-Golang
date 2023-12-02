package articlesdto

type CreateArticleRequest struct {
	ArticleName string `json:"articleName" form:"articleName" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Image       string `json:"image" form:"image"`
	UserID      int    `json:"userId" form:"userId"`
}

type UpdateArticleRequest struct {
	ArticleName string `json:"articleName" form:"articleName"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
	UserID      int    `json:"userId" form:"userId"`
}

type ArticleResponse struct {
	ID          int    `json:"id"`
	ArticleName string `json:"articleName"`
	Description string `json:"description"`
	Image       string `json:"image"`
	UserID      int    `json:"userId"`
}
