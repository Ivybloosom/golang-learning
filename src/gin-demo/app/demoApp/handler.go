/*
* @Author: Ivy
* @Date: 2022/5/26 12:20 PM
*
* 功能说明
**/
package demoApp

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// gin.Context，封装了request和response
func getting(c *gin.Context) {
	//c.String(http.StatusOK, "get: hello World!")
	c.JSON(http.StatusOK, gin.H{"message": "hello world"})
}

func getUser(c *gin.Context) {
	name := c.Param("name")
	info := c.Param("information")
	info = strings.Trim(info, "/")
	c.String(http.StatusOK, name+" : "+info)
}

func getUser2(c *gin.Context) {
	id := c.Query("id")
	c.DefaultQuery("id", "0")
	c.String(http.StatusOK, fmt.Sprintf("Your id is %s.", id))

}

// 定义接收数据的结构体
type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func login(c *gin.Context) {
	//接收 JSON 数据:  ShouldBindJSON()
	// 声明接收的变量
	var json Login
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&json); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 判断用户名密码是否正确
	if json.User != "root" || json.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})

	//接收表单数据: Bind()默认解析并绑定form格式
	//var form Login
	//if err := c.Bind(&form); err != nil {
	//	// 返回错误信息
	//	// gin.H封装了生成json数据的工具
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//if form.User != "root" || form.Password != "admin" {
	//	c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
	//	return
	//}
	//c.JSON(http.StatusOK, gin.H{"status": "200"})

	//接收uri数据
	//var login Login
	//// Bind()默认解析并绑定form格式
	//// 根据请求头中content-type自动推断
	//if err := c.ShouldBindUri(&login); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//// 判断用户名密码是否正确
	//if login.User != "root" || login.Password != "admin" {
	//	c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
	//	return
	//}
	//c.JSON(http.StatusOK, gin.H{"status": "200"})

}
