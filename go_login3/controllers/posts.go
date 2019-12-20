package controllers

import (
	"github.com/gin-gonic/gin"
	"go_learn/models"
	"net/http"
)

func GetPosts(c *gin.Context) {
	posts := models.GetPosts()
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":	  posts,
	})
}