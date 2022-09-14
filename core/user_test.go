package core

import (
	"fmt"
	"testing"

	"github.com/silverswords/sand/model"
	"github.com/silverswords/sand/services"
)

var instance *Application

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		"root", "my-123456", "mqying.xyz", "3306", "mall", "utf8", true, "Local")

	config := &Config{
		Dsn: dsn,
	}

	instance = CreateApplication(config)
	usersService := services.CreateUsersService(instance)
	productsService := services.CreateProductsService(instance)
	categoryService := services.CreateCategoryService(instance)
	ordersService := services.CreateOrdersService(instance)
	orderDetailsService := services.CreateOrderDetailsService(instance)
	shoppingCartsService := services.CreateShoppingCartsService(instance)
	service := services.CreateService(usersService, productsService, categoryService,
		ordersService, orderDetailsService, shoppingCartsService)
	instance.SetServices(&service)
}

func TestCreateUser(t *testing.T) {
	var user = model.User{UnionID: "1111", OpenID: "11111", Mobile: "12345678901"}
	if err := instance.Services().Users().Create(&user); err != nil {
		t.Errorf("CreateUser: %v", err)
	}
}

func TestUpdateMobile(t *testing.T) {
	var user = model.User{UnionID: "1111", OpenID: "1111", Mobile: "10987654321"}
	if err := instance.Services().Users().UpdateMobile(&user); err != nil {
		t.Errorf("CreateUser: %v", err)
	}
}
