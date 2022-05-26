package sort

func QuickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	pivot := list[0]
	var less []int
	var greater []int

	for i := range list {
		if list[i] <= pivot {
			less = append(less, i)
		} else {
			greater = append(greater, i)
		}
	}

	newList := make([]int, len(list), len(list))
	newList = append(newList, QuickSort(less)...)
	newList = append(newList, pivot)
	newList = append(newList, QuickSort(greater)...)

	return newList
}
