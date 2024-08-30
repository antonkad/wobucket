package main

import (
	"wobucket-server/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Register routes
	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
