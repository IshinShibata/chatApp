package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/IshinShibata/chatApp/ent"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	DbSetUp()

	client := GetClient()
	fmt.Println(client)

	SetupRouter().Run(":8080")
}

var (
	client *ent.Client
)

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

func GetClient() *ent.Client {
	return client
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/users", UsersLIst)
	//r.GET("/users/:id", UserShow)
	//r.POST("/users", UserCreate)
	//r.PUT("/users/:id", UserUpdate)
	//r.DELETE("/users/:id", UserDelete)
	return r
}

func UserCreate() string {
	return "test"
}

func UserShow() {

}

func UsersLIst(c *gin.Context) {
	users := GetClient().User.Query().AllX(context.Background())
	fmt.Println(users)
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func UserUpdate() {

}

func UserDelete() {

}
