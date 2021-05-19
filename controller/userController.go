package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"prometheus_devops/db"
	"prometheus_devops/model"
	"prometheus_devops/services"
	"strconv"
)

func Login(c *gin.Context) {
	name := c.PostForm("username")
	password := c.PostForm("password")

	user, err := services.GetUserByName(name)

	if err != nil {
		c.JSON(200, gin.H{"msg": fmt.Sprintf("Login failed %s", err)})
		return
	}
	if user.Password == password {
		c.JSON(200, gin.H{"msg": "Login success"})
		return
	}
	c.JSON(401, gin.H{"msg": "Login faied"})
}

func AddUser(c *gin.Context) {

	user := &model.User{}
	c.ShouldBind(&user)
	row := db.DB.Create(&user).RowsAffected
	if row == 1 {
		c.JSON(200, gin.H{"msg": "success"})
		return
	}
	c.JSON(200, gin.H{"msg": "failed create user"})
}

func GetUser(c *gin.Context) {
	id := c.Query("id")
	if id != "" {
		user, row := services.GetUser(id)
		if row == 1 {
			c.JSON(200, gin.H{"data": gin.H{"user": user}})

			return
		}
		c.JSON(200, gin.H{"msg": "failed find user"})
		return
	}
	user := services.GetAllUser()

	c.JSON(200, gin.H{"data": gin.H{"user": user}})
}

func DeleteUser(c *gin.Context) {
	id := c.Query("id")
	if services.DelteUser(id) {
		c.JSON(200, gin.H{"msg": "success"})
		return
	}
	c.JSON(200, gin.H{"msg": "failed delete user"})
}

func UpdateUser(c *gin.Context) {
	id := c.PostForm("id")
	username := c.PostForm("username")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	n, _ := strconv.Atoi(id)
	user := &model.User{
		Model:    gorm.Model{ID: uint(n)},
		Name:     username,
		Password: password,
		Phone:    phone,
	}
	if services.UpdateUser(user) {
		c.JSON(200, gin.H{"msg": "success"})
		return
	}
	c.JSON(200, gin.H{"msg": " update failed "})
}
