//Solution 1
// O(n^2) time | O(1) space
package main

func TwoNumberSum(array []int, target int) []int {
	for i := 0; i < len(array)-1; i++{
		firstNum := array[i]
		for j := i + 1; j < len(array); j++{
			secondNum := array[j]
			if firstNum+secondNum == target{
				return []int{firstNum, secondNum}
			}
		}
	}
	return []int{}
}

//Solution 2
// O(n) time | O(n) space
package main


func TwoNumberSum(array []int, target int) []int{
	nums := map[int]bool{}
	for _, num := range array{
		potentialMatch := target - num
		if _, found := nums[potentialMatch]; found{
			return []int{potentialMatch, num}
		}
		nums[num] = true
	}
	return []int{}
}

//Solution 3
//O(nlog(n)) | O(1) space
package main

import "sort"

func TwoNumberSum(array []int, target int) []int{
	sort.Ints(array)
	left, right := 0, len(array)-1
	for left < right {
		currentSum := array[left] + array[right]
		if currentSum == target{
			return []int{array[left], array[right]}
			// moving the either left or right
		} else if currentSum < target{
			left += 1
		} else {
			right -= 1
		}
	}
	return []int{}
}

// https://levelup.gitconnected.com/two-number-sum-in-golang-355627d6c861

// https://towardsdatascience.com/leetcode-two-sum-with-go-67d24e5a53f3