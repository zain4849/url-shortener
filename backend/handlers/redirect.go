package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HandleRedirect(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	var url URL
	shortURL := "http://localhost:8080/" + id
	db.First(&url, "short = ?", shortURL)
	return c.Redirect(http.StatusFound, url.Original)
}
