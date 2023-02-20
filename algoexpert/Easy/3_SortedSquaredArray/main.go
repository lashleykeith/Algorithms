//Solution 1
// O(nlogn) time | O(n)space
package main

import (
	"sort"
)

func SortedSquaredArray(array []int) []int {
	sortedSquares := make([]int, len(array))

	for idx, value := range array{
		sortedSquares[idx] = value * value
	}

	sort.Ints(sortedSquares)
	return sortedSquares
}

//Solution 2
// O(n) time | O(n)space
package main



func SortedSquaredArray(array []int) []int{
	sortedSquares := make([]int, len(array))

	smallerValueIdx := 0
	largerValueIdx := len(array) - 1

	for idx := len(array) - 1; idx >= 0; idx-- {
		smallerValue := array[smallerValueIdx]
		largerValue := array[largerValueIdx]

		if abs(smallerValue) > abs(largerValue) {
			sortedSquares[idx] = smallerValue * smallerValue
			smallerValueIdx += 1
		} else {
			sortedSquares[idx] = largerValue * largerValue
			largerValueIdx -= 1
		}
	}

	return sortedSquares
}

func abs(a int) int{
	if a < 0 {
		return -a
	}
	return a
}
