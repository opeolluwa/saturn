package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/opeolluwa/saturn/controllers"
)

func LoadRoutes(e *echo.Echo) *echo.Echo {
	e.GET("/health", controllers.HealthCheck)
	return e
}
