// O(nlogn) time | O(1) space

package main

import (
    "sort"
)


func MinimumWaitingTime(queries []int) int{
   sort.Ints(queries)

   totalWaitingTime := 0
   for idx, duration := range queries {
    queriesLeft := len(queries) - (idx + 1)
    totalWaitingTime += duration * queriesLeft
   }

   return totalWaitingTime
}

/*
We need to edit this code to print out the the totalWaitingTime of this query.

This is the query
{
  "queries": [3, 2, 1, 2, 6]

}

This is the code:
*/