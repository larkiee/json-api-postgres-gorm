package main

import (
	"fmt"

	globals "example.com/larkiee/interview/Globals"
	"example.com/larkiee/interview/handlesrs/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main(){
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	user.Register(e)
	
	e.Start(fmt.Sprintf(":%s", globals.SERVICE_PORT))
}