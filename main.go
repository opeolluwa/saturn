package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/opeolluwa/saturn/routers"
)

func main() {
	e := echo.New()

	app := routers.LoadRoutes(e)
	app.Use(middleware.CORS())
	app.Use(middleware.RequestLogger())
	
	e.Logger.Fatal(app.Start(":3345"))
}
