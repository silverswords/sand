package main

import (
	"fmt"
	"log"
	"os"

	json "github.com/json-iterator/go"
	"github.com/silverswords/sand/sql/mysql"
)

func main() {
	config := mysql.Config{
		Host:     "mqying.xyz",
		Port:     "3306",
		UserName: "root",
		Password: "123456",
		Database: "test",
		Charset:  "utf8",
	}

	data, err := json.Marshal(config)
	if err != nil {
		fmt.Println("error:", err)
	}

	err = os.WriteFile("sql", data, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
