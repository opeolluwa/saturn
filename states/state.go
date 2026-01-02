package states

import (
	"github.com/labstack/echo/v4"
	"github.com/opeolluwa/saturn/config"
	"gorm.io/gorm"
)

type State struct {
	Database *gorm.DB
	App      echo.Echo
}

func Init(environment config.Environment) (State, error) {

	db, err := config.ConnectDatabase(environment)
	if err != nil {
		return State{}, err
	}

	return State{
		Database: db,
		App:      *echo.New(),
	}, nil
}
