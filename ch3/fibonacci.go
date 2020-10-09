package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	result := 0
	for i := 0; i <= 10; i++ {
		result = fibbonacci(i)
		fmt.Printf("fibbonacci(%d)-----> result: %d \n",i, result)
	}
	end := time.Now()
	
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)

}

func fibbonacci(i int) (res int) {
	if i <= 1 {
		res = 1
	} else {
		res = fibbonacci(i - 1) + fibbonacci(i - 2)
	}
	return
}