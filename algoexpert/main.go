// # O(nlogn) time | O(1) space
package main

import (
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


/* 

this is the array of the value of coins we have
coins = [5,7,1,1,2,3,22]

We need to print the minimal amount that we can not create in the resulting variable currentChangeCreated with the following code.
 */