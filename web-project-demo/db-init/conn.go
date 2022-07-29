// 连接数据库
/*
* @Author: Ivy
* @Date: 2022/5/26 12:46 PM
 */
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB //声明一个全局的 db 变量

// 初始化 MySQL 函数
func InitMySQL() (err error) {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/go-curd?charset=utf8&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	return
}

func main() {
	// 初始化 MySQL
	err := InitMySQL()
	if err != nil {
		panic(err)
	}
	fmt.Println("db connected success!")
	defer db.Close()

	//db.Query()
	//显示全部数据
	//queryRowAllDemo()

}
