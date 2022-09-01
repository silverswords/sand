package services

import (
	"github.com/silverswords/sand/models/mysql"
	"github.com/silverswords/sand/models/structs"
)

func CreateOrderTable() error {
	if err := mysql.CreateOrderTable(); err != nil {
		return err
	}

	return nil
}

func InsertOrder(order structs.Order) error {
	if err := mysql.InsertOrder(order); err != nil {
		return err
	}

	return nil
}

func GetOrderBrifeInfoByOpenID(openID string) ([]*structs.Order, error) {
	brifeInfo, err := mysql.GetOrderBrifeInfoByOpenID(openID)
	if err != nil {
		return nil, err
	}

	return brifeInfo, nil
}

func GetOrderBrifeInfoByStoreID(storeID string) ([]*structs.Order, error) {
	brifeInfo, err := mysql.GetOrderBrifeInfoByStoreID(storeID)
	if err != nil {
		return nil, err
	}

	return brifeInfo, nil
}

func GetOrderDetialByOrderID(orderID string) (*structs.Order, error) {
	detial, err := mysql.GetOrderDetialByOrderID(orderID)
	if err != nil {
		return nil, err
	}

	return detial, nil
}

func ModifyOrderStatus(orderID string, status uint8) error {
	if err := mysql.ModifyOrderStatus(orderID, status); err != nil {
		return err
	}

	return nil
}
