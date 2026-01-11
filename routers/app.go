package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/opeolluwa/saturn/controllers"
	"github.com/opeolluwa/saturn/states"
)

func LoadRoutes(e *echo.Echo, state *states.State) *echo.Echo {
	loadAuthRoutes(e, state)
	
	e.GET("/health", controllers.HealthCheck)

	return e
}
