/*
* @Author: Ivy
* @Date: 2022/4/26 4:48 PM
**/
package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("nums = %v, target = %v\n", nums, target)
	answer := twoSum(nums, target)
	fmt.Println(answer)
}

func twoSum(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
