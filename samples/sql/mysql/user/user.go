package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	mysql "github.com/silverswords/sand/models/mysql"
	"github.com/silverswords/sand/models/structs"
)

func main() {
	user := structs.User{
		UnionId: "UnionId_test",
		OpenId:  "OpenId_test",
		Mobile:  "123456",
	}

	if err := mysql.CreateTable(); err != nil {
		fmt.Println(err)
	}

	if err := mysql.CreateUser(user.UnionId, user.OpenId); err != nil {
		fmt.Println(err)
	}

	if err := mysql.ModifyMobile(user.UnionId, time.Now(), user.Mobile); err != nil {
		fmt.Println(err)
	}

	userInfo, err := mysql.GetUserInfo(user.UnionId)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Mobile: %s", userInfo.Mobile)
}
