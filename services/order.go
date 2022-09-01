package services

import (
	"github.com/silverswords/sand/models/mysql"
)

func CreateOrderTable() error {
	if err := mysql.CreateOrderTable(); err != nil {
		return err
	}

	return nil
}
