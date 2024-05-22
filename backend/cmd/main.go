package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	dbConn "github.com/user123/URL-shortener-Golang/backend/db"
	"github.com/user123/URL-shortener-Golang/backend/handlers"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{echo.GET, echo.POST},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Connect to DB
	db, err := dbConn.ConnectToDB()
	if err != nil {
		panic("Unable to connect to the database")
	}

	db.Exec("DELETE FROM urls")

	// Apply schema
	db.AutoMigrate(&handlers.URL{})

	// **************** Routes ****************
	e.POST("/shorten", func(c echo.Context) error {
		return handlers.HandleShorten(c, db)
	})

	// Redirect
	e.GET("/:id", func(c echo.Context) error {
		return handlers.HandleRedirect(c, db)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
