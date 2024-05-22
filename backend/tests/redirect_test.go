package tests

import (
	"net/http"
	"net/http/httptest" // To create mock http request and responses
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/user123/URL-shortener-Golang/backend/handlers"
)

func TestHandleRedirect(t *testing.T) {
	// Set up the test database
	db, teardown := SetupTestDB()
	defer teardown()

	// Insert a test URL record into the database
	testURL := handlers.URL{Original: "http://example.com", Short: "http://localhost:8080/test"}
	if err := db.Create(&testURL).Error; err != nil {
		t.Fatalf("Failed to insert test URL record into the database: %v", err)
	}

	// Create a new Echo instance
	e := echo.New()

	// Create a new request to simulate a redirect
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("test")

	// Call the HandleRedirect function
	err := handlers.HandleRedirect(c, db)
	if err != nil {
		t.Fatalf("HandleRedirect function failed: %v", err)
	}

	// Assert the redirect status code
	assert.Equal(t, http.StatusFound, rec.Code)

	// Assert the Location header to match the original URL
	assert.Equal(t, testURL.Original, rec.Header().Get("Location"))
}
