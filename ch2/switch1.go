package main
import "fmt"

func main() {
	var num1 int = 100
	switch num1 {
	case 98, 99:
		fmt.Println("It`s equal to 98")
		fallthrough
	case 100:
		fmt.Println("It`s equal to 98")
		fallthrough
	default:
		fmt.Println("It's not equal to 98 or 100")
	}
}