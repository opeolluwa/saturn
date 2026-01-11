package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/opeolluwa/saturn/controllers"
	"github.com/opeolluwa/saturn/services"
	"github.com/opeolluwa/saturn/states"
)

func loadAuthRoutes(state *states.State) *echo.Echo {
	e := &state.App

	authService := services.NewAuthService(state.Database)
	authController := controllers.NewAuthController(authService)

	authRoutes := e.Group("/auth")
	authRoutes.POST("/signup", authController.SignUp)

	return e
}
