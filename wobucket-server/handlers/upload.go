package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

const uploadDir = "./uploads"

func HandleUploadBlob(c echo.Context) error {

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid file upload"})
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to open file"})
	}
	defer src.Close()

	// Ensure the upload directory exists
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to create upload directory"})
	}

	// Destination file path
	dstPath := filepath.Join(uploadDir, file.Filename)

	// Create the destination file
	dst, err := os.Create(dstPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to create destination file"})
	}
	defer dst.Close()

	// Copy the uploaded file to the destination
	if _, err = dst.ReadFrom(src); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to save file"})
	}

	// Respond with a success message
	return c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully", file.Filename))
}
