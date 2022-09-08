package core

import (
	"github.com/silverswords/sand/services"
	"gorm.io/gorm"
)

type Application struct {
	config  *Config
	gormDB  *mysqlDatabase
	service services.Service
}

func CreateApplication(config *Config) *Application {
	app := &Application{
		gormDB: NewMysqlDatabase("mysql", config.Dsn),
	}

	return app
}

func (a *Application) GetDefaultGormDB() *gorm.DB {
	return a.gormDB.DB
}

func (a *Application) SetServices(s *services.Service) {
	a.service = *s
}

func (a *Application) Services() services.Service {
	return a.service
}
