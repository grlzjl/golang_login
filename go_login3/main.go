package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"go_learn/controllers"
	"go_learn/database"
	"net/http"
)

func main() {
	r := gin.Default()

	database.InitDB()

	//加载路由
	r.GET("/ping", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"Status":	0,
			"data": 	"Hello World!",
		})
	})

	r.POST("/login", controllers.Login)

	r.GET("/article", controllers.GetPosts)

	// 设置静态文件夹
	r.Static("/static", "./static")


	r.Run(":8083")
}
