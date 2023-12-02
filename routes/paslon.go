package routes

import (
	"Restful_Go/handlers"
	"Restful_Go/pkg/middleware"
	"Restful_Go/pkg/postgres"
	"Restful_Go/repositories"

	"github.com/labstack/echo/v4"
)

func PaslonRoutes(e *echo.Group) {
	r := repositories.RepositoryPaslon(postgres.DB)
	h := handlers.PaslonHandler(r)

	e.GET("/paslons", h.FindPaslons)
	e.GET("/paslon/:id", h.GetPaslon)
	e.POST("/paslon", middleware.UploadFile(h.CreatePaslon))
	e.PATCH("/paslon/:id", middleware.UploadFile(h.UpdatePaslon))
	e.DELETE("/paslon/:id", h.DeletePaslon)
}
