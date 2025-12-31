package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/opeolluwa/saturn/controllers"
)

func Routes(e echo.Echo) {
	e.GET("/health", controllers.HealthCheck)
}
