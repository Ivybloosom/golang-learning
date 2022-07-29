package main

import (
	"fmt"
	"go-learning/algorithm/search/search"
)

var (
	a        = make([]int, 1024*1024, 1024*1024)
	findItem int
)

func main() {

	for i := 0; i < 1024*1024; i++ {
		a[i] = i
	}
	//fmt.Println(a)

	fmt.Println("请输入要查找的数据：")
	fmt.Scanln(&findItem)

	index := search.BinarySearch(a, findItem)
	if index != -1 {
		fmt.Printf("已查找到数据 %v, 所在索引为 %v。\n", findItem, index)
	} else {
		fmt.Println("没有找到数据")
	}

}
