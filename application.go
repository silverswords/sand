package sand

import (
	"fmt"

	"github.com/silverswords/sand/core"
	"github.com/silverswords/sand/services"
)

var Application *core.Application

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		"root", "my-123456", "mqying.xyz", "3306", "mall", "utf8", true, "Local")

	config := &core.Config{
		Dsn: dsn,
	}

	Application = core.CreateApplication(config)
	usersService := services.CreateUsersService(Application)
	productsService := services.CreateProductsService(Application)
	categoryService := services.CreateCategoryService(Application)
	ordersService := services.CreateOrdersService(Application)
	orderDetailsService := services.CreateOrderDetailsService(Application)
	shoppingCartsService := services.CreateShoppingCartsService(Application)
	virtualStore := services.CreateVirtualStoreService(Application)
	service := services.CreateService(usersService, productsService, categoryService,
		ordersService, orderDetailsService, shoppingCartsService, virtualStore)
	Application.SetServices(&service)
}
