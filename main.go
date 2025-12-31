package main

import (
	"github.com/labstack/echo/v4"
	"github.com/opeolluwa/saturn/routers"
)

func main() {
	e := echo.New()

	app := routers.LoadRoutes(e)

	e.Logger.Fatal(app.Start(":3345"))
}
