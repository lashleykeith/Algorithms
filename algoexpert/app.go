package main 

import (
    "fmt"
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

func main() {
    array := []int{3, 4, 5, 6, 7, 8}
    sortedSquares := SortedSquaredArray(array)

    for _, value := range sortedSquares {
        fmt.Println(value)
    }
}
