package routes

import (
	"Restful_Go/handlers"
	"Restful_Go/pkg/middleware"
	"Restful_Go/pkg/postgres"
	"Restful_Go/repositories"

	"github.com/labstack/echo/v4"
)

func ArticleRoutes(e *echo.Group) {
	r := repositories.RepositoryArticle(postgres.DB)
	h := handlers.ArticleHandler(r)

	e.GET("/articles", h.FindArticles)
	e.GET("/article/:id", h.GetArticle)
	e.POST("/article", middleware.Auth(middleware.UploadFile(h.CreateArticle)))
	e.PATCH("/article/:id", middleware.Auth(middleware.UploadFile(h.UpdateArticle)))
	e.DELETE("/article/:id", h.DeleteArticle)
}
