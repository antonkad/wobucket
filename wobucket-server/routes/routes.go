package routes

import (
	"wobucket-server/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/upload", handlers.HandleUploadBlob)
}
