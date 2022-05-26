/*
* @Author: Ivy
* @Date: 2022/5/25 12:15 AM
*
* 功能说明
**/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fileName := "src/my_gopl/read_file/test.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	b2 := bufio.NewReader(os.Stdin)
	s2, _ := b2.ReadString('\n')
	fmt.Println(s2)
}
