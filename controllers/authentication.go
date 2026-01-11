package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/opeolluwa/saturn/adapters/requests"
	"github.com/opeolluwa/saturn/adapters/responses"
	"github.com/opeolluwa/saturn/services"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return AuthController{authService}
}

func (c *AuthController) SignUp(ctx echo.Context) error {

	var createUserRequest requests.CreateUserRequest
	if err := ctx.Bind(&createUserRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.Error("badly formatted request"))
	}

	if err := ctx.Validate(&createUserRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.Error(err.Error()))
	}

	if err := c.authService.CreateUser(ctx, createUserRequest); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, responses.New("Account created successfully"))
}
