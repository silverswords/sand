package model

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var (
// 	username = "root"
// 	password = "my-123456"
// 	host     = "mqying.xyz"
// 	port     = 3306
// 	Dbname   = "product"

// 	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
// 		username, password, host, port, Dbname)
// )

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Log("Failed to connect to the database, error: " + err.Error())
	}
	db.AutoMigrate(Product{})

	id, err := CreateProduct(db, 0, 1, 66.88, `[
		{
		"category_id":-7923833148285436,
		"name":"有",
		"status":710049414037452,
		"created_at":"1974-03-14 23:29:06"
		},
		{
		"category_id":-6952204478771452,
		"name":"按",
		"status":807399053007360,
		"created_at":"1977-04-13 02:54:19"
		},
		{
		"category_id":5316619045525976,
		"name":"进",
		"status":5755071993643720,
		"created_at":"2011-10-16 09:01:53"
		}
		]`, "main-title", "subtitle", 0, 888)
	if err != nil {
		t.Log(err)
	}

	t.Log(id)
}

func TestListAllProducts(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Log("Failed to connect to the database, error: " + err.Error())
	}
	db.AutoMigrate(Product{})

	products, err := ListAllProducts(db)
	if err != nil {
		t.Log(err)
	}

	for _, v := range products {
		t.Log(v)
	}
}

func TestModifyProduct(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Log("Failed to connect to the database, error: " + err.Error())
	}
	db.AutoMigrate(Product{})

	ModifyProduct(db, 4, "store_id", 8)
	ModifyProduct(db, 4, "category_id", 8)
	ModifyProduct(db, 4, "price", 88.88)
	ModifyProduct(db, 4, "photo_urls", "[]")
	ModifyProduct(db, 4, "main_title", "main")
	ModifyProduct(db, 4, "subtitle", "sub")
	ModifyProduct(db, 4, "status", 1)
	ModifyProduct(db, 4, "stock", 999)
}

func TestQueryByProductID(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Log("Failed to connect to the database, error: " + err.Error())
	}
	db.AutoMigrate(Product{})

	product, err := QueryByProductId(db, 1)
	if err != nil {
		t.Log(err)
	}

	t.Log(product)
}

func TestQueryByStoreID(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Log("Failed to connect to the database, error: " + err.Error())
	}
	db.AutoMigrate(Product{})

	product, err := QueryByStoreId(db, 12)
	if err != nil {
		t.Log(err)
	}
	t.Log(len(product))
	for _, v := range product {
		t.Log(v)
	}
}

func TestDeleteByProductID(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Log("Failed to connect to the database, error: " + err.Error())
	}
	db.AutoMigrate(Product{})

	err = DeleteByProductID(db, 11)
	if err != nil {
		t.Log(err)
	}
}

func TestDeleteByStoreID(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Log("Failed to connect to the database, error: " + err.Error())
	}
	db.AutoMigrate(Product{})

	err = DeleteByStoreID(db, 1)
	if err != nil {
		t.Log(err)
	}
}
