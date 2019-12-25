// package main

// import (
// 	"database/sql"

// 	"github.com/adefemi171/newVirel/handlers"

// 	"github.com/labstack/echo"
// 	"github.com/labstack/echo/middleware"
// 	_ "github.com/mattn/go-sqlite3"
// )

// func main() {
// 	e := echo.New()

// 	// Middleware
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	//Initialize database
// 	db := initDB("storage.db")
// 	migrate(db)

// 	// Define the HTTP routes
// 	e.File("/", "public/index.html")
// 	e.GET("/polls", handlers.GetPolls(db))
// 	e.PUT("/poll/:index", handlers.UpdatePoll(db))

// 	// e.POST("/polls/", func(c echo.Context) error {
// 	// 	return c.JSON(200, "POST Success")
// 	// })
// 	// e.GET("/polls", func(c echo.Context) error {
// 	// 	return c.JSON(200, "GET Success")
// 	// })

// 	// e.PUT("/polls", func(c echo.Context) error {
// 	// 	return c.JSON(200, "PUT Success")
// 	// })

// 	// e.PUT("/polls/:id", func(c echo.Context) error {
// 	// 	return c.JSON(200, "UPDATE Success "+" "+c.Param("id"))
// 	// })

// 	// Start server
// 	e.Logger.Fatal(e.Start(":9000"))
// }

// func initDB(filepath string) *sql.DB {
// 	db, err := sql.Open("sqlite3", filepath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if db == nil {
// 		panic("Database Nil")
// 	}
// 	return db
// }

// func migrate(db *sql.DB) {
// 	sql := `
// 		CREATE TABLE IF NOT EXISTS polls(
// 			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
// 			f_name VARCHAR NOT NULL,
// 			m_name  VARCHAR,
// 			l_name VARCHAR NOT NULL,
// 			img  VARCHAR NOT NULL,
// 			upvotes INTEGER NOT NULL,
// 			downvotes INTEGER NOT NULL,
// 			UNIQUE(f_name, m_name, l_name)

// 		);

// 		INSERT OR IGNORE INTO polls(f_name, m_name, l_name, img, upvotes, downvotes) VALUES('Adefemi', 'Micheal', 'Afuwape', './img/adefemi.png', 1, 0);
// 		INSERT OR IGNORE INTO polls(f_name, m_name, l_name, img, upvotes, downvotes) VALUES('Bayonle', ' ', 'Adeyemi', './img/adefemi.png', 1, 0);
// 	`
// 	_, err := db.Exec(sql)

// 	if err != nil {
// 		panic(err)
// 	}
// }

package main

import (
	// "database/sql"
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

	// Initialize the database
	db := initDB("storage.db")
	migrate(db)

	// Define the HTTP routes
	e.File("/", "public/index.html")
	e.GET("/polls", handlers.GetPolls(db))
	e.PUT("/poll/:index", handlers.UpdatePoll(db))

	// e.GET("/polls", func(c echo.Context) error {
	// 	return c.JSON(200, "GET Polls")
	// })

	// e.PUT("/polls", func(c echo.Context) error {
	// 	return c.JSON(200, "PUT Polls")
	// })

	// e.PUT("/polls/:id", func(c echo.Context) error {
	// 	return c.JSON(200, "UPDATE Poll "+c.Param("id"))
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
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
            CREATE TABLE IF NOT EXISTS polls(
                    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                    name VARCHAR NOT NULL,
                    topic VARCHAR NOT NULL,
                    src VARCHAR NOT NULL,
                    upvotes INTEGER NOT NULL,
                    downvotes INTEGER NOT NULL,
                    UNIQUE(name)
            );

            INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Potus','Adefemi Micheal', 'https://cdn.colorlib.com/wp/wp-content/uploads/sites/2/angular-logo.png', 1, 0);

            INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Dharmani', 'Oyogho Daniel','https://upload.wikimedia.org/wikipedia/commons/thumb/5/53/Vue.js_Logo.svg/400px-Vue.js_Logo.svg.png', 1, 0);

            INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Ray','Raven Holmes','https://upload.wikimedia.org/wikipedia/commons/thumb/a/a7/React-icon.svg/1200px-React-icon.svg.png', 1, 0);

            INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('FirstLady','Hawa Kromah','https://cdn-images-1.medium.com/max/741/1*9oD6P0dEfPYp3Vkk2UTzCg.png', 1, 0);

            INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Victoria','Victoria Lucas','https://images.g2crowd.com/uploads/product/image/social_landscape/social_landscape_1489710848/knockout-js.png', 1, 0);
       `
	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}
