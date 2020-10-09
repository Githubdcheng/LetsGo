package main

import "fmt"

func main() {
	var min, max int
	min, max = Minmax(78, 69)
	fmt.Printf("Minmium is %d, Maxmium is %d \n", min, max)
}

func Minmax(x, y int) (min, max int) {
	if x < y {
		min, max = x, y
	} else {
		min, max = y, x
	}
	return
}