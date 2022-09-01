package services

import (
	"github.com/silverswords/sand/models/mysql"
)

func CreateUserTable() error {
	if err := mysql.CreateUserTable(); err != nil {
		return err
	}

	return nil
}
