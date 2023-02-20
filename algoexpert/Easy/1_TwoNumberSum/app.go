package main

import "fmt"

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

func main() {
    arr := []int{-3, 4, 5, 3, 7, 8, 3, 2}
    targetSum := 10

    result := TwoNumberSum(arr, targetSum)

    if len(result) == 2 {
        fmt.Printf("The two numbers that add up to %d are %d and %d\n", targetSum, result[0], result[1])
    } else {
        fmt.Printf("No two numbers in the array add up to %d\n", targetSum)
    }
}
