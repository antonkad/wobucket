package routes

import (
	"wobucket-server/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.POST("/upload", handlers.HandleUploadBlob)
	e.GET("/files", handlers.HandleListFiles)
	e.DELETE("/files/:name", handlers.HandleDeleteFile)
}
