package main

import "fmt"

func ValidSubsequence(array []int, sequence []int) ([]int, bool) {
    subsequence := []int{}
    arrIdx := 0
    seqIdx := 0
    for arrIdx < len(array) && seqIdx < len(sequence) {
        if array[arrIdx] == sequence[seqIdx] {
            subsequence = append(subsequence, array[arrIdx])
            seqIdx += 1
        }
        arrIdx += 1
    }
    return subsequence, seqIdx == len(sequence)
}

func main(){
array := []int{2, 3, 4, 5, 6}
sequence := []int{2, 4, 5}
subsequence, isValid := ValidSubsequence(array, sequence)
    if isValid {
        fmt.Println(subsequence) // Output: [2 4 5]
    }
}
