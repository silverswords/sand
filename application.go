package sand

import (
	"fmt"
	"io/ioutil"

	json "github.com/json-iterator/go"
	"github.com/silverswords/sand/core"
	"github.com/silverswords/sand/services"
)

var application *core.Application

type Config struct {
	Host     string
	Port     string
	Database string
	UserName string
	Password string
	Charset  string
}

func init() {
	c := &Config{}

	data, err := ioutil.ReadFile("../config/sql.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(data, c); err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		c.UserName, c.Password, c.Host, c.Port, c.Database, c.Charset, true, "Local")

	config := &core.Config{
		Dsn: dsn,
	}

	application = core.CreateApplication(config)

	usersService := services.CreateUsersService(application)
	productsService := services.CreateProductsService(application)
	categoryService := services.CreateCategoryService(application)
	ordersService := services.CreateOrdersService(application)
	shoppingCartsService := services.CreateCartsService(application)
	virtualStore := services.CreateVirtualStoreService(application)
	weChat := services.CreateWeChatService()
	sign := services.CreateSignService()
	if err != nil {
		panic(err)
	}

	service := services.CreateService(usersService, productsService, categoryService,
		ordersService, shoppingCartsService, virtualStore, weChat, sign)
	application.SetServices(&service)
}

func GetApplication() *core.Application {
	return application
}
