package main

import (
    "fmt"
    "sort"
)

func MinimumWaitingTime(queries []int) int {
    sort.Ints(queries)

    totalWaitingTime := 0
    for idx, duration := range queries {
        queriesLeft := len(queries) - (idx + 1)
        totalWaitingTime += duration * queriesLeft
    }

    fmt.Println("Total waiting time:", totalWaitingTime)
    return totalWaitingTime
}

func main() {
    queries := []int{3, 2, 1, 2, 6}
    MinimumWaitingTime(queries)
}
