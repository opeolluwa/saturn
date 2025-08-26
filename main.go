package main

import (
	"github.com/labstack/echo/v4"
	"github.com/opeolluwa/saturn/router"
)

func main() {
	e := echo.New()
	router.LoadRoutes(e)
}
