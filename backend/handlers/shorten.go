package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// URL represents the URL entity
type URL struct {
	ID       uint   `gorm:"primary_key"`
	Original string `validate:"required"`
	Short    string
}

// HandleShorten handles the /shorten endpoint
// HandleShorten handles the /shorten endpoint
func HandleShorten(c echo.Context, db *gorm.DB) error {
	// Define a struct to capture incoming request data
	type requestData struct {
		URL string `json:"url"`
	}

	// Initialize an instance of the struct to store the parsed request data
	var reqData requestData

	// Bind the request body to the requestData struct
	if err := c.Bind(&reqData); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request data")
	}

	// Extract the URL from the requestData struct
	original := reqData.URL

	// Check if the URL is empty
	if original == "" {
		return c.String(http.StatusBadRequest, "URL cannot be empty")
	}

	// Generate a short URL
	shortURL := Shorten(original)

	url := URL{
		Original: original,
		Short:    shortURL,
	}

	// Create new URL record in the database
	if err := db.Create(&url).Error; err != nil {
		return err
	}

	// Return the short URL as JSON object
	return c.JSONPretty(http.StatusOK, map[string]string{"shortenedUrl": shortURL}, "  ")
}

// Generate a short URL
func Shorten(url string) string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	length := 8

	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		randomString[i] = alphabet[rng.Intn(len(alphabet))]
	}

	shortURL := fmt.Sprintf("http://localhost:8080/%s", randomString)
	return shortURL
}
