package sort

func SelectSort(array []int) {
	var tmp int
	for i := range array {
		tmp = array[i]
		for j := i + 1; j < len(array); j++ {
			if tmp < array[j] {
				tmp, array[j] = array[j], tmp
			}
			array[i] = tmp
		}
	}
}
