package main

import "fmt"

type SpecialArray []interface{}

func ProductSum(array SpecialArray) int {
    return productSumHelper(array, 1)
}

func productSumHelper(array SpecialArray, depth int) int {
    sum := 0
    for _, v := range array {
        switch t := v.(type) {
        case SpecialArray:
            sum += productSumHelper(t, depth+1)
        default:
            sum += v.(int)
        }
    }
    return sum * depth
}

func main() {
    array := SpecialArray{5, 2, SpecialArray{7, -1}, 3, SpecialArray{6, SpecialArray{-13, 8}, 4}}
    result := ProductSum(array)
    fmt.Println(result) // Output should be 12
}
