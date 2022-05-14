package main

import (
	"fmt"
)

func main() {
	a := []int{10, 3, 5, 12, 9, 20}
	fmt.Println("input: ", a)
	sorted := false
	numArr := len(a) - 1
	for !sorted {
		sorted = true
		for i := 0; i < numArr; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				sorted = false
			}
		}
		numArr--
	}
	fmt.Println("sort: ", a)
}
