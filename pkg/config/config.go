package config

import (
	"context"
	"fmt"
	"github.com/IshinShibata/chatApp/ent"
	"log"
	"os"
)

var client *ent.Client

func DbSetUp() {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DATABASE")
	connection := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, pass, host, dbname)

	var err error
	client, err = ent.Open("mysql", connection)

	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
