package routes

import (
	"Restful_Go/handlers"
	"Restful_Go/pkg/postgres"
	"Restful_Go/repositories"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {
	r := repositories.RepositoryAuth(postgres.DB)
	h := handlers.AuthHandler(r)

	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
}
