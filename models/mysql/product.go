package mysql

import (
	"fmt"

	"github.com/silverswords/sand/models/structs"
)

const ProductTableName = "product"

const (
	mysqlCreateProductTable = iota
	mysqlInsertProduct
	mysqlProductBrifeInfo
	mysqlProductdetialInfo
	mysqlProductOfVirtualStore
)

var (
	productSQLString = []string{
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			pro_id 			VARCHAR(128) UNIQUE NOT NULL,
			store_id		VARCHAR(128) UNIQUE NOT NULL,
			price			DOUBLE NOT NULL DEFAULT 9999.99,
			main_title		VARCHAR(100) NOT NULL DEFAULT " ",
			subtitle		VARCHAR(100) NOT NULL DEFAULT " ",
			images			JSON,
			stock			INT UNSIGNED NOT NULL DEFAULT 0,
			status			INT UNSIGNED NOT NULL DEFAULT 0,
			created_time  	DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (pro_id)
		)  ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;`, ProductTableName),
		fmt.Sprintf(`INSERT INTO %s (pro_id, store_id, price, main_title, subtitle, images, stock, status) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, ProductTableName),
		fmt.Sprintf(`SELECT pro_id, price, subtitle, images, status FROM %s`, ProductTableName),
		fmt.Sprintf(`SELECT pro_id, price, main_title, subtitle, images, stock, status, created_time FROM %s WHERE pro_id = ?`, ProductTableName),
		fmt.Sprintf(`SELECT pro_id, price, main_title, subtitle, images, stock, status, created_time FROM %s WHERE store_id = ?`, ProductTableName),
	}
)

// CreateTable create product table
func CreateProductTable() error {
	_, err := db.Exec(productSQLString[mysqlCreateProductTable])
	if err != nil {
		return err
	}

	return nil
}

// InsertProduct insert product into the table
func InsertProduct(product structs.Product) error {
	_, err := db.Exec(productSQLString[mysqlInsertProduct], product.ProID,
		product.StoreID, product.Price, product.MainTitle, product.Subtitle,
		product.Images, product.Stock, product.Status)
	if err != nil {
		return err
	}

	return nil
}

// GetAll get all brife products info to homepage
func GetAllProduce() ([]*structs.Product, error) {
	rows, err := db.Query(productSQLString[mysqlProductBrifeInfo])
	if err != nil {
		return nil, err
	}

	var result []*structs.Product
	for rows.Next() {
		var (
			pro_id   string
			price    float64
			subtitle string
			images   interface{}
			status   uint8
		)
		if err := rows.Scan(&pro_id, &price, &subtitle, &images, &status); err != nil {
			return nil, err
		}

		result = append(result, &structs.Product{
			ProID:    pro_id,
			Price:    price,
			Subtitle: subtitle,
			Images:   images,
			Status:   status,
		})
	}

	return result, nil
}

// GetProductByID detial info in product page, got by id
func GetProductInfoByID(productID string) (*structs.Product, error) {
	var (
		pro_id       string
		price        float64
		main_title   string
		subtitle     string
		images       interface{}
		stock        uint64
		status       uint8
		created_time string

		result *structs.Product
	)

	err := db.QueryRow(productSQLString[mysqlProductdetialInfo], productID).Scan(&pro_id, &price, &main_title,
		&subtitle, &images, &stock, &status, &created_time)
	if err != nil {
		return nil, err
	}

	result = &structs.Product{
		ProID:      pro_id,
		Price:      price,
		MainTitle:  main_title,
		Subtitle:   subtitle,
		Images:     images,
		Stock:      stock,
		Status:     status,
		CreateTime: created_time,
	}

	return result, nil
}

// VirtualStoreProduct get virtual store's products by storeID
func GetVirtualStoreProsByID(storeID string) ([]*structs.Product, error) {
	rows, err := db.Query(productSQLString[mysqlProductOfVirtualStore], storeID)
	if err != nil {
		return nil, err
	}

	var result []*structs.Product
	for rows.Next() {
		var (
			pro_id       string
			price        float64
			main_title   string
			subtitle     string
			images       interface{}
			stock        uint64
			status       uint8
			created_time string
		)
		if err := rows.Scan(&pro_id, &price, &main_title, &subtitle, &images, &stock, &status, &created_time); err != nil {
			return nil, err
		}

		result = append(result, &structs.Product{
			ProID:      pro_id,
			Price:      price,
			MainTitle:  main_title,
			Subtitle:   subtitle,
			Images:     images,
			Stock:      stock,
			Status:     status,
			CreateTime: created_time,
		})
	}

	return result, nil
}
