// File: shorten_test.go
package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/user123/URL-shortener-Golang/backend/handlers"
)

func TestHandleShorten(t *testing.T) {
	db, teardown := SetupTestDB()
	defer teardown()
	e := echo.New()

	t.Run("Valid URL", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/shorten", strings.NewReader(`{"url": "http://example.com"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, handlers.HandleShorten(c, db)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), "http://localhost:8080/")
		}
	})

	t.Run("Very Long URL", func(t *testing.T) {
		longURL := "http://" + strings.Repeat("a", 2000) + ".com"
		req := httptest.NewRequest(http.MethodPost, "/shorten", strings.NewReader(`{"url": "`+longURL+`"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, handlers.HandleShorten(c, db)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), "http://localhost:8080/")
		}
	})

	t.Run("Empty URL", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/shorten", strings.NewReader(`{"url": ""}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, handlers.HandleShorten(c, db)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "URL cannot be empty", rec.Body.String())
		}
	})

}
