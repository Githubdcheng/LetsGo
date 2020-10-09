package main

import "fmt"

func main() {
	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", Mul3results(2, 5, 6))
}

func Mul3results(a, b, c int) int {
	return a * b * c
}
