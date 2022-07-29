// 定义 apps 路由信息
/*
* @Author: Ivy
* @Date: 2022/5/26 12:20 PM
 */

package demoApp

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {

	// 2.绑定路由规则，执行的函数
	// 路由组 为了管理一些相同的URL
	v1 := router.Group("/test")
	{
		v1.GET("/", getting)
		v1.GET("/user/:name/*information", getUser)
		v1.GET("/user", getUser2)
		v1.POST("/login", login)
	}

	v2 := router.Group("/release")
	{
		v2.GET("/", getting)
		v2.GET("/user/:name/*information", getUser)
		v2.GET("/user", getUser2)
		v2.POST("/", login)
	}

	//router.PUT("/", putting)
	//router.DELETE("/", deleting)
	//router.PATCH("/", patching)
	//router.HEAD("/", head)
	//router.OPTIONS("/", options)
}
