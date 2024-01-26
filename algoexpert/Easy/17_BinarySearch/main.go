package main

import "fmt"

func BinarySearch(array []int, target int) int {
	min := 0
	max := len(array) - 1
	for min <= max {
		avg := (max + min) / 2
		fmt.Printf("min: %d, max: %d, avg: %d\n", min, max, avg)
		if target > array[avg] {
			min = avg + 1
		} else if target < array[avg] {
			max = avg - 1
		} else {
			fmt.Printf("Element %d found at index %d\n", target, avg)
			return avg
		}
	}
	fmt.Printf("Element %d not found in the array\n", target)
	return -1
}

func main() {
	array := []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}
	target := 33

	result := BinarySearch(array, target)
	fmt.Printf("Search result: %d\n", result)
}
