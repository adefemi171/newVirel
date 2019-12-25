// package handlers

// import (
// 	"database/sql"
// 	"net/http"
// 	"strconv"

// 	"github.com/adefemi171/newVirel/models"
// 	"github.com/labstack/echo"
// )

// //Mapping strings as keys and anything else as values
// type H map[string]interface{}

// func GetPolls(db *sql.DB) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		return c.JSON(http.StatusOK, models.GetPolls(db))
// 	}
// }

// func UpdatePoll(db *sql.DB) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var poll models.Poll
// 		c.Bind(&poll)
// 		index, _ := strconv.Atoi(c.Param("index"))
// 		mat_no, err := models.UpdatePoll(db, index, poll.MatNo, poll.FName, poll.MName, poll.LName, poll.Img, poll.Upvotes, poll.Downvotes)
// 		if err == nil {
// 			return c.JSON(http.StatusCreated, H{
// 				"affected": mat_no,
// 			})
// 		}
// 		return err
// 	}
// }


 package handlers

    import (
        "database/sql"
        "net/http"
        "strconv"
		"github.com/labstack/echo"
		"github.com/adefemi171/newVirel/models"
	)

	type H map[string]interface{}

	func GetPolls(db *sql.DB) echo.HandlerFunc {
        return func(c echo.Context) error {
            return c.JSON(http.StatusOK, models.GetPolls(db))
        }
    }

    func UpdatePoll(db *sql.DB) echo.HandlerFunc {
        return func(c echo.Context) error {
            var poll models.Poll

            c.Bind(&poll)

            index, _ := strconv.Atoi(c.Param("index"))

            id, err := models.UpdatePoll(db, index, poll.Name, poll.Upvotes, poll.Downvotes)

            if err == nil {
                return c.JSON(http.StatusCreated, H{
                    "affected": id,
                })
            }

            return err
        }
    }