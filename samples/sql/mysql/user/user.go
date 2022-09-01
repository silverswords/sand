package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	mysql "github.com/silverswords/sand/models/mysql"
	"github.com/silverswords/sand/models/structs"
)

func main() {
	user := structs.User{
		UnionID: "UnionID_test",
		OpenID:  "OpenID_test",
		Mobile:  "123456",
	}

	if err := mysql.CreateUserTable(); err != nil {
		fmt.Println(err)
	}

	if err := mysql.InsertUser(user); err != nil {
		fmt.Println(err)
	}
}
