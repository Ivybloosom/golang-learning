/*
* @Author: Ivy
* @Date: 2022/5/26 9:53 AM
 */
package main

import (
	"fmt"

	"example.com/server/apps/app-demo"
	"example.com/server/router-init"
)

func main() {
	// 1.创建路由引擎
	//router := gin.Default()
	//router.SetTrustedProxies([]string{"127.0.0.1"})
	// 2.加载路由
	routers.Include(demoApp.Router)
	router := routers.Init()
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	err := router.Run(":8000")
	if err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
