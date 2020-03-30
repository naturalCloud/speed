package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
}

func (bs *Controller) success(c *gin.Context, data map[string]interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": "10000",
		"msg":  "ok",
		"data": data,
	})
}

func (bs *Controller) fail(c *gin.Context, errCode int, msg string) {

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		`code`: errCode,
		`msg`:  msg,
		`data`: gin.H{},
	})
}

// f 格式化的字符串
// s 输出的字符串
func (bs *Controller) TextRes(c *gin.Context, f string, s ...string) {
	c.String(http.StatusOK, f, s)

}
