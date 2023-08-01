package main

import "fmt"

func GetNthFib(n int) int {
	if n == 1 {
		return 0
	}
	num1, num2 := 0, 1
	for i := 3; i <= n; i++ {
		num1, num2 = num2, num2+num1
		if num2 == 34 {
			fmt.Printf("Fibonacci number at position %d is %d\n", i, num2)
		}
	}
	return num2
}

func main() {
	GetNthFib(34)
}
