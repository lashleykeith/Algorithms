package main

import "fmt"

func TwoNumberSum( array []int, target int) []int{
	for i := 0; i < len(array)-1; i++{
		firstNum := array[i]
		for j := 0; j < len(array); j++{
			secondNum := array[j]
			if firstNum + secondNum == target{
				return []int{firstNum, secondNum}
			}
		}
	}
	return []int{}
}

func main(){

	arr := []int{-3, 4, 5, 3, 7, 8, 3, 2}
	targetSum := 10

	result := TwoNumberSum{arr, targetSum}
}

// TwoNumberSum

/* 
For i as long as i is less 
that the array minus one we are 
going to go up.

Then we will have a first Number to the the array of i
second number to array of j

After that we need to see if the first and second 
number add up to the target

then return this these two numbers

return the []int{}

make an array -3, 4, 5, 3, 7, 8, 3, 2 arr
give this a target of 10 targetSum

place the result
*/