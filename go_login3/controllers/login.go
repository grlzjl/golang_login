package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type User struct {
	Username string `json:"userName"`
	Password string `json:"passWord"`
}

//GetLabels 获取所有的标签
func Login(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	var u User
	json.Unmarshal(data, &u)

	var result string

	if u.Username == "zhangsan" && u.Password == "123" {
		result = "ok"

	} else {
		result = "err"
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   result,
	})

}
