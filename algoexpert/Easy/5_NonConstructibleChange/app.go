package main

import (
    "fmt"
    "sort"
)

func NonConstructibleChange(coins []int) int{
    sort.Ints(coins)

    currentChangeCreated := 0

    for _, coin := range coins{
        if coin > currentChangeCreated + 1 {
            return currentChangeCreated + 1
        }
        currentChangeCreated += coin
    }
    return currentChangeCreated + 1
}

func main() {
    coins := []int{5,7,1,1,2,3,22}
    fmt.Println(NonConstructibleChange(coins)) // Output: 21
}
