package main

import (
	"fmt"
	"go-learning/algorithm/sort/sort"
)

func main() {
	a := make([]int, 20, 20)
	for i := range a {
		if i < len(a)/2 {
			a[i] = len(a) - i*2 - 1
		} else {
			a[i] = 2*len(a) - i*2
		}
	}
	fmt.Printf(" original array: %v\n", a)

	// 选择排序  O(n^2)  不稳定
	a1 := a
	go sort.SelectSort(a1)
	fmt.Printf(" `select_sorted` array: %v\n", a1)

	// 插入排序
	a2 := a
	go sort.InsertSort(a2)
	fmt.Printf(" insert_sorted array: %v\n", a2)

	// 快速排序
	a3 := a
	go sort.QuickSort(a3)
	fmt.Printf(" quick_sorted array: %v\n", a3)

}
