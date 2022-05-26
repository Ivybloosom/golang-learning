/*
* @Author: Ivy
* @Date: 2022/5/25 6:41 PM
*
* go-mysql curd demo
**/
package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB //声明一个全局的 db 变量

// 定义接收数据的结构
type user struct {
	id      int
	name    string
	phone   int
	email   string
	addTime time.Time
}

// 初始化 MySQL 函数
func initMySQL() (err error) {
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

// 查找全部数据
func queryRowAllDemo() []user {
	// 声明查询语句
	sqlStr := "SELECT * FROM user"

	var uList []user
	// 执行查询并且扫描至 u
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println("查询多条数据错误", err)
	}

	for rows.Next() {
		var u user
		rows.Scan(&u.id, &u.name, &u.phone, &u.email, &u.addTime)
		//加入数组
		uList = append(uList, u)
	}

	return uList

	//fmt.Println("================= 已查到全部数据 =================")
	//for _, user := range uList {
	//	fmt.Printf("id:%d  name:%s  phone:%d  email:%s  addTime:%v\n", user.id, user.name, user.phone, user.email, user.addTime)
	//}
	//
	//fmt.Println()
}

// 查找一行数据
func queryRowDemo(id int, name string, phone int, email string) {
	// 声明查询语句
	sqlStr := "SELECT id,name,phone,email,addTime FROM user WHERE id = ? OR name = ? OR phone = ? OR email = ?"
	// 声明一个 user 类型的变量
	var u user
	// 执行查询并且扫描至 u
	err := db.QueryRow(sqlStr, id, name, phone, email).Scan(&u.id, &u.name, &u.phone, &u.email, &u.addTime)
	if err != nil {
		fmt.Println("查询数据错误", err)
		return
	}
	fmt.Println("================= 已查到数据 =================")
	fmt.Printf("id:%d  name:%s  phone:%d  email:%s  addTime:%v\n", u.id, u.name, u.phone, u.email, u.addTime)

}

// 增加一行数据
func insertRowDemo(name string, phone int, email string) {
	sqlStr := "INSERT INTO user(name, phone, email, addTime) VALUES(?, ?, ?, Now())"
	result, err := db.Exec(sqlStr, name, phone, email)
	if err != nil {
		fmt.Printf("insert data failed, err:%v\n", err)
		return
	}
	newID, err := result.LastInsertId() //新增数据的ID
	if err != nil {
		fmt.Printf("get insert lastInsertId failed, err:%v\n", err)
		return
	}
	n, _ := result.RowsAffected() //受影响行数
	fmt.Println("================= 增加数据 =================")
	fmt.Printf("insert success!\nid: %d, affected rows: %d \n", newID, n)
}

// 删除一行数据
func deleteRowDemo(id int) {
	sqlStr := "DELETE FROM user WHERE id = ?"
	result, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete data failed, err:%d\n", err)
		return
	}
	n, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get affected failed, err:%v\n", err)
		return
	}
	fmt.Println("================= 删除数据 =================")
	fmt.Printf("delete success!\naffected rows:%d\n", n)
}

// 更新一组数据
func updateRowDemo(id int, name string, phone int, email string) {
	sqlStr := "UPDATE user SET name = ?, phone = ?, email = ? WHERE id = ?"
	result, err := db.Exec(sqlStr, name, phone, email, id)
	if err != nil {
		fmt.Printf("update data failed, err:%v\n", err)
		return
	}
	n, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get rowsaffected failed, err:%v\n", err)
		return
	}
	fmt.Println("================= 更新数据 =================")
	fmt.Printf("update success!\naffected rows:%d\n", n)
}

func main() {
	// 初始化 MySQL
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	fmt.Println("db connected success!")
	defer db.Close()

	//初始化路由
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	//显示全部数据
	var uList []user
	uList = queryRowAllDemo()
	var result string
	result = result + "================= 已查到全部数据 =================\n"
	//for _, user := range uList {
	//	result += fmt.Sprintf("id:%d  name:%s  phone:%d  email:%s  addTime:%v\n", user.id, user.name, user.phone, user.email, user.addTime)
	//}
	var resultList []string
	for _, user := range uList {
		resultList = append(resultList, fmt.Sprintf("id:%d  name:%s  phone:%d  email:%s  addTime:%v", user.id, user.name, user.phone, user.email, user.addTime))
	}
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, result)
		c.JSON(http.StatusOK, gin.H{"result": resultList})
	})

	// 查找数据
	//queryRowDemo(2, "", 0, "")

	// 增加数据
	//insertRowDemo("liang", 1736103002, "zhao@gmail.com")
	//queryRowAllDemo()

	// 更新数据
	//updateRowDemo(4, "xianxue", 1320000000, "xian@qq.com")
	//queryRowAllDemo()

	// 删除数据
	//deleteRowDemo(4)
	//queryRowAllDemo()

	// 监听端口设置在 8000
	if err := router.Run(":8000"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
