package search

import "fmt"

func BinarySearch(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		fmt.Println(mid)
		if list[mid] == item {
			return mid
		} else if list[mid] < item {
			low = mid + 1
		} else if list[mid] > item {
			high = mid - 1
		}
	}

	return -1
}
