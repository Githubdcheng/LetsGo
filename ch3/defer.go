package main
import (
	"fmt"
	"time"
)

func main () {
	func1()
}

func func1() {
	fmt.Printf("in func 1 at the top \n")
	defer func2()
	fmt.Printf("in func 1 at the bottom \n")
}

func func2() {
	
	fmt.Printf("in func 2 at Deferred until the end of the calling function!")
}