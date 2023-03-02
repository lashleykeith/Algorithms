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
    var n, targetSum int
    fmt.Print("Enter the number of elements in the array: ")
    fmt.Scan(&n)

    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Printf("Enter element #%d: ", i+1)
        fmt.Scan(&arr[i])
    }

    fmt.Print("Enter the target sum: ")
    fmt.Scan(&targetSum)

    result := TwoNumberSum(arr, targetSum)

    if len(result) == 2 {
        fmt.Printf("The two numbers that add up to %d are %d and %d\n", targetSum, result[0], result[1])
    } else {
        fmt.Printf("No two numbers in the array add up to %d\n", targetSum)
    }
}