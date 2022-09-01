package services

import (
	"github.com/silverswords/sand/models/mysql"
	"github.com/silverswords/sand/models/structs"
)

func CreateProductTable() error {
	if err := mysql.CreateProductTable(); err != nil {
		return nil
	}

	return nil
}

func InsertProduct(product structs.Product) error {
	if err := mysql.InsertProduct(product); err != nil {
		return err
	}

	return nil
}

func GetAllProduce() ([]*structs.Product, error) {
	result, err := mysql.GetAllProduce()
	if err != nil {
		return nil, err
	}

	return result, err
}

func GetProductInfoByID(productID string) (*structs.Product, error) {
	result, err := mysql.GetProductInfoByID(productID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetVirtualStoreProsByID(storeID string) ([]*structs.Product, error) {
	result, err := mysql.GetVirtualStoreProsByID(storeID)
	if err != nil {
		return nil, err
	}

	return result, nil
}