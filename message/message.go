package mes

import "github.com/gin-gonic/gin"

func NewRes(c *gin.Context, code int, httpStatus int, msg string, data interface{}) {
	c.JSON(code, gin.H{
		"code": httpStatus,
		"msg":  msg,
		"data": data,
	})
}

func Success(c *gin.Context, httpStatus int, data interface{}) {
	NewRes(c, 200, 200, "success", data)
}

func Fail(c *gin.Context, httpStatus int, msg string, err interface{}) {
	NewRes(c, 200, httpStatus, msg, err)
}
