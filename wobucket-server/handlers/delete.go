package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func HandleDeleteFile(c echo.Context) error {
	// Get the file name from the URL parameter
	fileName := c.Param("name")

	// Build the full file path
	filePath := filepath.Join(uploadDir, fileName)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "File not found"})
	}

	// Delete the file
	if err := os.Remove(filePath); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to delete file"})
	}

	// Return success message
	return c.JSON(http.StatusOK, map[string]string{"message": "File deleted successfully"})
}
