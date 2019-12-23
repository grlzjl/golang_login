package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

	fmt.Println(result)

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   result,
	})

}
