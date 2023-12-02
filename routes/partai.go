package routes

import (
	"Restful_Go/handlers"
	"Restful_Go/pkg/middleware"
	"Restful_Go/pkg/postgres"
	"Restful_Go/repositories"

	"github.com/labstack/echo/v4"
)

func PartaiRoutes(e *echo.Group) {
	r := repositories.RepositoryPartai(postgres.DB)
	h := handlers.PartaiHandler(r)

	e.GET("/partais", h.FindPartais)
	e.GET("/partai/:id", h.GetPartai)
	e.POST("/partai", middleware.UploadFile(h.CreatePartai))
	e.PATCH("/partai/:id", middleware.UploadFile(h.UpdatePartai))
	e.DELETE("/partai/:id", h.DeletePartai)
}
