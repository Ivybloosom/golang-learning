package main

import "fmt"

func insertionSort(a *[6]int) {
	var key int
	var i int

	for j := 1; j < 6; j++ {
		key = a[j]
		i = j - 1
		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i--
		}
		a[i+1] = key
	}
}

func main() {

	a := [6]int{8, 2, 4, 9, 3, 6}

	insertionSort(&a)

	fmt.Println(a)
}
