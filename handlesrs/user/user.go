package user

import (
	"net/http"

	userController "example.com/larkiee/interview/controllers/user"
	"github.com/labstack/echo/v4"
)

var userCtl = userController.New()

type errJson struct {
	Code int
	Msg string
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	res, err := userCtl.GetUser(id)
	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, errJson{
				Code: http.StatusNotFound,
				Msg: "not found user",
			})
			return err
		}
		c.JSON(http.StatusInternalServerError, errJson{
			Code: http.StatusInternalServerError,
			Msg: "server internal fault",
		})
		return err
	}
	c.JSON(http.StatusOK, res)
	return nil
}


func Register(e *echo.Echo) {
	e.GET("/user/:id", GetUser)
}