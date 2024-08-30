package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleUploadBlob(c echo.Context) error {

	return c.String(http.StatusOK, "File uploaded")
}
