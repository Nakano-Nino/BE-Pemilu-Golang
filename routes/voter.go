package routes

import (
	"Restful_Go/handlers"
	"Restful_Go/pkg/middleware"
	"Restful_Go/pkg/postgres"
	"Restful_Go/repositories"

	"github.com/labstack/echo/v4"
)

func VoterRoutes(e *echo.Group) {
	r := repositories.RepositoryVoter(postgres.DB)
	h := handlers.VoterHandler(r)

	e.GET("/voters", h.FindVoters)
	e.POST("/vote", middleware.Auth(h.Vote))
}
