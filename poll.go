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
	e.POST("/polls/", func(c echo.Context) error {
		return c.JSON(200, "POST Success")
	})
	e.GET("/polls", func(c echo.Context) error {
		return c.JSON(200, "GET Success")
	})

	e.PUT("/polls", func(c echo.Context) error {
		return c.JSON(200, "PUT Success")
	})

	e.PUT("/polls/:id", func(c echo.Context) error {
		return c.JSON(200, "UPDATE Success "+" "+c.Param("id"))
	})

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}
