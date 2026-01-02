package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/opeolluwa/saturn/controllers"
	"github.com/opeolluwa/saturn/states"
)

func LoadRoutes(state *states.State) *echo.Echo {
	e := &state.App
	
	e.GET("/health", controllers.HealthCheck)
	return e
}
