package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/opeolluwa/saturn/adapters/requests"
	"github.com/opeolluwa/saturn/repositories"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(db *gorm.DB) AuthService {
	userRepository := repositories.NewUserRepository(db)
	return AuthService{userRepository}
}

func (s *AuthService) CreateUser(ctx echo.Context, req requests.CreateUserRequest) error {

	err := s.userRepository.CreateUser(req)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return nil
}
