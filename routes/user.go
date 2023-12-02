package routes

import (
	"Restful_Go/handlers"
	"Restful_Go/pkg/postgres"
	"Restful_Go/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	r := repositories.RepositoryUser(postgres.DB)
	h := handlers.UserHandler(r)

	e.GET("/users", h.FindUsers)
	e.GET("/user/:id", h.GetUsers)
	e.POST("/user", h.CreateUser)
}
