package main

import (
	"database/sql"

	"github.com/adefemi171/newVirel/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Initialize database
	db := initDB("storage.db")
	migrate(db)

	// Define the HTTP routes
	e.File("/", "public/index.html")
	e.GET("/polls", handlers.GetPolls(db))
	e.PUT("/poll/:index", handlers.UpdatePoll(db))

	// e.POST("/polls/", func(c echo.Context) error {
	// 	return c.JSON(200, "POST Success")
	// })
	// e.GET("/polls", func(c echo.Context) error {
	// 	return c.JSON(200, "GET Success")
	// })

	// e.PUT("/polls", func(c echo.Context) error {
	// 	return c.JSON(200, "PUT Success")
	// })

	// e.PUT("/polls/:id", func(c echo.Context) error {
	// 	return c.JSON(200, "UPDATE Success "+" "+c.Param("id"))
	// })

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("Database Nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
		CREATE TABLE IF NOT EXISTS polls(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			f_name VARCHAR NOT NULL,
			m_name  VARCHAR,
			l_name VARCHAR NOT NULL,
			img  VARCHAR NOT NULL,
			upvotes INTEGER NOT NULL,
			downvotes INTEGER NOT NULL,
			UNIQUE(f_name, m_name, l_name)

		);

		INSERT OR IGNORE INTO polls(f_name, m_name, l_name, img, upvotes, downvotes) VALUES('Adefemi', 'Micheal', 'Afuwape', './img/adefemi.png', 1, 0);
		INSERT OR IGNORE INTO polls(f_name, m_name, l_name, img, upvotes, downvotes) VALUES('Bayonle', ' ', 'Adeyemi', './img/adefemi.png', 1, 0);
	`
	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}
