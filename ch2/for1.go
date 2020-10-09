package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}

	for i, j := 0, 5; i < j; i, j = i+1, j-1 {
		fmt.Printf("This is the %d %d iteration\n", i, j)
	}
}