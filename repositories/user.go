package repositories

import (
	"strings"

	"github.com/google/uuid"
	"github.com/opeolluwa/saturn/models"
	"github.com/opeolluwa/saturn/adapters/requests"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db}
}

func (r *UserRepository) CreateUser(req requests.CreateUserRequest) error {
	identifier, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(req.Password)), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		FirstName:  strings.TrimSpace(req.FistName),
		LastName:   strings.TrimSpace(req.LastName),
		Email:      strings.ToLower(strings.TrimSpace(req.Email)),
		Identifier: identifier,
		Password:   string(password),
	}

	result := r.db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
