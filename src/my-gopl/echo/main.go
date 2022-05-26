/*
* @Author: Ivy
* @Date: 2022/5/24 7:57 PM
*
* echo 输出命令行参数
**/
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行参数如下:")
	for _, text := range os.Args[1:] {
		fmt.Printf("%v ", text)

	}

	fmt.Println()
}
