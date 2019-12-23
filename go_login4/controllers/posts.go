package controllers

import (
	"log"
	"net/http"
	"strconv"

	"fmt"

	"my-blog-by-go/models"

	"github.com/gin-gonic/gin"
)

//GetPosts 获取所有的文章
func GetPosts(c *gin.Context) {
	posts := models.GetPosts()
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   posts,
	})
}

//GetPostByLabelId 根据label id获取post
func GetPostByLabelId(c *gin.Context) {
	labelid := c.Param("labelid")
	if labelid == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 1,
			"msg":    "labelid 不能为空",
		})
	}
	labelId, err := strconv.ParseInt(labelid, 10, 64)
	postIDs, err := models.GetPostsByLabelID(labelId)
	if err != nil {
		log.Println(labelId, err)
		c.JSON(http.StatusNotFound, gin.H{
			"status": 1,
			"msg":    "GetPostsByLabelID error, err:" + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   models.GetPostByIDs(postIDs),
	})
}

//GetPostByCategoryId category id获取post
func GetPostByCategoryId(c *gin.Context) {
	categoryid := c.Param("categoryid")
	if categoryid == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 1,
			"msg":    "categoryid 不能为空",
		})
	}
	categoryId, err := strconv.ParseInt(categoryid, 10, 64)
	postIDs, err := models.GetPostsByCategoryId(categoryId)
	if err != nil {
		log.Println(categoryId, err)
		c.JSON(http.StatusNotFound, gin.H{
			"status": 1,
			"msg":    "GetPostsByCategoryId error, error:" + err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   models.GetPostByIDs(postIDs),
	})
}

//GetPostByCategoryId category id获取post
func GetPostById(c *gin.Context) {
	postid := c.Param("postid")
	if postid == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 1,
			"msg":    "postid 不能为空",
		})
		return
	}
	postId, err := strconv.ParseInt(postid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "postid 格式错误",
		})
		return
	}

	postInfo := models.GetPostByID(postId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "GetPostByID err:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   postInfo,
	})
}

//GetPosts 获取所有的文章
func InsertPosts(c *gin.Context) {

	requestInfo := models.InsertPostReq{}
	err := c.BindJSON(&requestInfo)
	if err != nil {
		fmt.Println("Parse body failed. Error:", err)
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "params 格式错误",
		})
		return
	}

	// 检查文章是否已经存在
	postInfo := models.GetPostByTitle(requestInfo.Title)
	if postInfo != nil && postInfo.Id > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": 2,
			"msg":    "title 已存在",
		})
		return
	}

	id, err := models.InsertPostByTransaction(requestInfo)
	if err != nil {
		log.Println("InsertPostByTransaction failed. Error:", err)
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "transaction error",
		})
		return
	}

	posts := models.GetPosts()
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"data":    posts,
		"post_id": id,
		"msg":     "success",
	})
}

//LikePosts 文章点赞
func LikePosts(c *gin.Context) {

	postid := c.Param("postid")
	if postid == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 1,
			"msg":    "postid 不能为空",
		})
		return
	}

	postId, err := strconv.ParseInt(postid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "postid 格式错误",
		})
		return
	}

	err = models.AddLikeByPostId(postId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "系统错误, err:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
	})
	return
}
