package user

import (
	"context"
	"github.com/IshinShibata/chatApp/ent/user"
	"github.com/IshinShibata/chatApp/pkg/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserCreate(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name is required",
		})
		return
	}
	user, err := mysql.GetClient().User.Create().SetName(name).Save(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSONP(http.StatusOK, gin.H{
		"message": "created user",
		"user":    user,
	})
}

func UserShow(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	user, err := mysql.GetClient().User.Query().Where(user.ID(id)).Only(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve user",
		})
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"message": "user detail",
		"user":    user,
	})
}

func UsersLIst(c *gin.Context) {
	users := mysql.GetClient().User.Query().AllX(context.Background())
	c.JSON(http.StatusOK, gin.H{
		"message": "user list",
		"users":   users,
	})
}

func UserUpdate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}
	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name is required",
		})
		return
	}
	user := mysql.GetClient().User.UpdateOneID(id).SetName(name)
	if err := user.Exec(context.Background()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user updated",
	})
}

func UserDelete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}
	user := mysql.GetClient().User.DeleteOneID(id)
	if err := user.Exec(context.Background()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSONP(http.StatusOK, gin.H{
		"message": "user deleted",
	})
}
