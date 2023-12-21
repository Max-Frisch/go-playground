package main

import "fmt"

func main() {
	var num int = 42
	var name string = "bob"
	var isTrue bool = true
	var riches float64 = 555.55
	fmt.Printf("variable num: %d is of type %T\n", num, num)
	fmt.Printf("variable name: %s is of type %T\n", name, name)
	fmt.Printf("variable isTrue: %v is of type %T\n", isTrue, isTrue)
	fmt.Printf("variable riches: %.2f is of type %T\n", riches, riches)
}
