package main

import (
    "sort"
)


func TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int{
    sort.Ints(redShirtSpeeds)
    sort.Ints(blueShirtSpeeds)

    if !fastest{
        reverseArrayInPlace(redShirtSpeeds)
    }

    totalSpeed := 0
    for idx := range redShirtSpeeds{
        rider1 := redShirtSpeeds[idx]
        rider2 := blueShirtSpeeds[len(blueShirtSpeeds)-idx-1]
        totalSpeed += max(rider1, rider2)
    }
    return totalSpeed

}

func reverseArrayInPlace(array []int){
    var start = 0
    var end = len(array) - 1
    for start < end {
        temp := array[start]
        array[start] = array[end]
        array[end] = temp
        start += 1
        end -= 1
    }
}


func max(a, b int) int{
    if a > b {
        return a
    }
    return b
}

/*
create a function TandemBicycle holding in the parameters redShirtSpeeds,  blueShirtSpeeds and a fastest for the bool

sort the red and blue ShirtSpeeds

if not the fastest we need to reverseArrayInPlace

create the reverseArrayInPace function it hols an array
  intialize start
  intialize end to be at the end of the array

  for loop as long as start is less than end
*/