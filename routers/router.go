package routers

import (
	"github.com/gin-gonic/gin"
	"prometheus_devops/controller"
	"prometheus_devops/middleware"
)

func SerRouter(mode string) *gin.Engine {
	gin.SetMode(mode)
	r := gin.Default()
	r.Use(middleware.Cors)
	r.POST("/login", controller.Login)
	r.POST("/addUser", controller.AddUser)
	r.GET("/getUser", controller.GetUser)
	r.GET("/delUser", controller.DeleteUser)
	r.POST("/updateUser", controller.UpdateUser)

	return r
}
