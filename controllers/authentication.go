package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// authService := service.NewAuthService()

func SignupUser(c echo.Context) error {
	return c.String(http.StatusOK, "Service is healthy")
}
