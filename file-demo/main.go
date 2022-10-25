// 功能说明
/*
* @Author: Ivy
* @Date: 2022/9/21 3:53 PM
 */
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	companyListResp := make([]string, 5)
	companyListResp, err := getCompanyNameList()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(companyListResp)
}

func getCompanyNameList() ([]string, error) {
	companyNameList := make([]string, 0, 5)
	companyNameListInput := make([]string, 0, 300)

	// 读取文件
	file, err := os.Open("file-demo/公司名称.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	for {
		// 读到结尾退出
		if !scanner.Scan() {
			break
		}

		// 读取行
		line := scanner.Text()

		// 跳过空行
		if line == "" {
			continue
		}

		// 去掉首尾空格换行等
		line = strings.TrimSpace(line)
		companyNameListInput = append(companyNameListInput, line)

		// 分割字符串
		//lineSlice := strings.Split(line, " ")
	}

	n := len(companyNameListInput)

	// 生成随机的返回数组
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r := rand.Intn(n)
		companyNameList = append(companyNameList, companyNameListInput[r])
	}

	return companyNameList, nil
}
