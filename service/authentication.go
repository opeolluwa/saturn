package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/opeolluwa/saturn/repositories"
	"github.com/opeolluwa/saturn/requests"
	"gorm.io/gorm"
)

type AuthService struct {
	Repository repositories.UserRepository
}

func NewAuthService(db *gorm.DB) AuthService {
	repo := repositories.NewUserRepository(db)
	return AuthService{Repository: repo}
}

func (s *AuthService) CreateUser(ctx echo.Context, req requests.CreateUserRequest) error {

	err := s.Repository.CreateUser(req)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}
