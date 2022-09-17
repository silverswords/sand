package model

import (
	"fmt"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	username = "root"
	password = "my-123456"
	host     = "mqying.xyz"
	port     = 3306
	Dbname   = "product"

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port, Dbname)
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Log("Failed to connect to the database, error: " + err.Error())
	}
	db.AutoMigrate(Product{})

	for i := 0; i < 5; i++ {
		id, err := CreateProduct(db, 1, 2, 66.88, `[
			{
			"category_id":-7923833148285436,
			"name":"有",
			"status":710049414037452,
			"created_at":"1974-03-14 23:29:06"
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

func TestModify(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Log("Failed to connect to the database, error: " + err.Error())
	}
	db.AutoMigrate(Product{})

	ModifyCategoryID(db, []uint64{2, 1}, 999)
	ModifyPhotoUrls(db, 1, "[]")
	ModifyMainTitle(db, 1, "main")
	ModifyStoreID(db, []uint64{1, 2}, 888)
	ModifySubtitle(db, 1, "sub")
	ModifyStatus(db, []uint64{5, 6}, 2)
	ModifyStock(db, []uint64{10, 11}, 6666)
	ModifyPrice(db, 1, 111.11)
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

func TestModifyProduct(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Log("Failed to connect to the database, error: " + err.Error())
	}
	db.AutoMigrate(Product{})

	product := Product{
		Model:   Model{ID: 1},
		StoreID: 999,
		Price:   8888.8,
	}

	ModifyProduct(db, &product)
}

func TestListByCategoryID(t *testing.T) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Log("Failed to connect to the database, error: " + err.Error())
	}
	db.AutoMigrate(Product{})

	result, err := ListByCategoryID(db, 999)
	if err != nil {
		t.Log(err)
	}

	for _, v := range result {
		t.Log(v)
	}

	t.Log(len(result))
}
