package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/opeolluwa/saturn/controllers"
	"github.com/opeolluwa/saturn/states"
)

func loadAuthRoutes(state *states.State) *echo.Echo {
	e := &state.App
	authRoutes := e.Group("/auth")
	
	authRoutes.POST("/signup", controllers.SignupUser)
	return e
}