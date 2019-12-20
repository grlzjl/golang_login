package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadRouters(router *gin.Engine) {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Status": 0,
			"data":   "Hello World!",
		})
	})
}
