package main

import (
	// "database/sql"

	// "database/sql"

	// "github.com/adefemi171/newVirel/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define the HTTP routes
	e.GET("/polls", func(c echo.Context) error {
		return c.JSON(200, "GET Polls")
	})

	e.PUT("/polls", func(c echo.Context) error {
		return c.JSON(200, "PUT Polls")
	})

	e.PUT("/polls/:id", func(c echo.Context) error {
		return c.JSON(200, "UPDATE Poll "+c.Param("id"))
	})

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}
