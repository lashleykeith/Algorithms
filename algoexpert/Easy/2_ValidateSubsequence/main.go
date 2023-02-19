//Solution 1

package main

// O(n) time | O(1) space
func IsValidSubsequence(array []int, sequence []int) bool{
	arrIdx := 0
	seqIdx := 0
	for arrIdx < len(array) && seqIdx < len(sequence){
		if array[arrIdx] == sequence[seqIdx]{
			seqIdx += 1
		}
		arrIdx += 1
	}
	return seqIdx == len(sequence)
}



//Solution 2
// O(n) time | O(1) space

func IsValidSubsequence(array []int, sequence []int) bool{
	seqIdx := 0
	for _, value := range array{
		if seqIdx == len(sequence){
			break
		}
		if value == sequence[seqIdx]{
			seqIdx += 1
		}
	}
	return seqIdx == len(sequence)
}


//https://cloudiosx.medium.com/leetcode-problem-2-validate-subsequence-34060d52b3e7