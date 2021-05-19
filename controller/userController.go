package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"prometheus_devops/db"
	mes "prometheus_devops/message"
	"prometheus_devops/model"
	"prometheus_devops/services"
	"strconv"
)

func Login(c *gin.Context) {
	//name := c.PostForm("username")
	//password := c.PostForm("password")
	var u model.User
	err2 := c.ShouldBind(&u)
	log.Println("err2 >>>>>>>:", err2)
	user, err := services.GetUserByName(u.Name)

	if err != nil {
		mes.Fail(c, http.StatusUnauthorized, "Login failed", err.Error())

		return
	}
	if user.Password == u.Password {
		mes.Success(c, http.StatusOK, gin.H{"token": "adjkfajdfghjadgfygadjbfhjvbewfjk"})

		return
	}

	mes.Fail(c, http.StatusUnauthorized, "Login failed", err.Error())
}

func AddUser(c *gin.Context) {

	user := &model.User{}
	c.ShouldBind(&user)
	row := db.DB.Create(&user).RowsAffected
	if row == 1 {
		mes.Success(c, http.StatusOK, nil)
		return
	}
	mes.Fail(c, http.StatusUnprocessableEntity, "failed create user", nil)

}

func GetUser(c *gin.Context) {
	id := c.Query("id")
	if id != "" {
		user, row := services.GetUser(id)
		if row == 1 {
			mes.Success(c, http.StatusOK, user)
			return
		}
		mes.Fail(c, http.StatusUnprocessableEntity, "failed find user", nil)

		return
	}
	user := services.GetAllUser()
	mes.Success(c, http.StatusOK, user)

}

func DeleteUser(c *gin.Context) {
	id := c.Query("id")
	if services.DelteUser(id) {
		mes.Success(c, 200, id)

		return
	}
	mes.Fail(c, http.StatusUnprocessableEntity, "failed delete user", nil)

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
		mes.Success(c, http.StatusOK, user)

		return
	}
	mes.Fail(c, http.StatusUnprocessableEntity, "update failed", nil)

}
