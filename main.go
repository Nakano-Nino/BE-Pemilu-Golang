package main

import (
	"Restful_Go/database"
	"Restful_Go/pkg/postgres"
	"Restful_Go/routes"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	postgres.DatabaseInit()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))

	fmt.Println("Server running on port 5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}
