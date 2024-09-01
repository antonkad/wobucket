package handlers

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func HandleListFiles(c echo.Context) error {

	// Open the uploads directory
	files, err := os.ReadDir(uploadDir)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to read upload directory"})
	}

	// Prepare a slice to hold file names
	var fileList []string

	// Loop through directory entries
	for _, file := range files {
		if !file.IsDir() {
			fileList = append(fileList, file.Name())
		}
	}

	// Return the list of files as a JSON response
	return c.JSON(http.StatusOK, fileList)
}
