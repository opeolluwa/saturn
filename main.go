package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/opeolluwa/saturn/router"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	router.LoadRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
