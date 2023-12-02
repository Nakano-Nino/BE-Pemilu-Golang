package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	UserRoutes(e)
	PaslonRoutes(e)
	ArticleRoutes(e)
	PartaiRoutes(e)
	AuthRoutes(e)
	VoterRoutes(e)
}
