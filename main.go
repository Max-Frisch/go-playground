package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "2222"
	var wurst bool = true
	i, err := strconv.Atoi(s)
	fmt.Printf("%d, %T\n", i, i)
	fmt.Printf("%t, %T\n", wurst, wurst)
	fmt.Printf("%v, %T\n", err, err)
}
