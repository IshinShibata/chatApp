package config

import (
	"github.com/IshinShibata/chatApp/internal/domain/user"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/users", user.UsersLIst)
	r.GET("/users/:id", user.UserShow)
	r.POST("/users", user.UserCreate)
	r.PUT("/users/:id", user.UserUpdate)
	r.DELETE("/users/:id", user.UserDelete)
	return r
}
