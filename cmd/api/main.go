package main

import (
	"fmt"
	"github.com/IshinShibata/chatApp/pkg/config"
	"github.com/IshinShibata/chatApp/pkg/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	mysql.DbSetUp()

	client := mysql.GetClient()
	fmt.Println(client)

	config.SetupRouter().Run(":8080")
}
