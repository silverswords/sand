package application

import (
	"github.com/silverswords/sand/services"
	"gorm.io/gorm"
)

type Application struct {
	gormDB *mysqlDatabase
}

type Config struct {
	Dsn string
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

func (a *Application) Services() services.Service {
	return nil
}
