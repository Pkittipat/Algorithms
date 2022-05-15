package main

import "fmt"

func main() {
	arr := []int{10, 5, 2, 7, 4, 9, 12, 1, 8, 6, 11, 3}
	fmt.Println("input: ", arr)
	mergeSort(arr)
	fmt.Println("output", arr)
}

func mergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	mergeSort(left)
	mergeSort(right)

	i, j := 0, 0
	tmp := []int{}
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			tmp = append(tmp, left[i])
			i++
		} else {
			tmp = append(tmp, right[j])
			j++
		}
	}

	tmp = append(tmp, left[i:]...)
	tmp = append(tmp, right[j:]...)
	for i, v := range tmp {
		arr[i] = v
	}
}
