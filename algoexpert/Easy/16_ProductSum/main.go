package main

import "fmt"

type SpecialArray []interface{}

func ProductSum(array []interface{}) int{
    fmt.Println(array[0])
    return productSumHelper(array, 1)
}


func productSumHelper(array SpecialArray, depth int) int{
    sum := 0
    for _, v := range array{
        switch t := v.(type){
        case SpecialArray:
            sum += productSumHelper(t, depth + 1)

        default:
            sum += v.(int)
        }
    }
    return sum * depth
}

// array = [5, 2, [7, -1], 3, [6, [-13, 8], 4]]