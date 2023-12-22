package main

import "fmt"

func main() {
	grades := [5]int{1, 2, 3, 4, 5}
	fmt.Println(grades)
	fmt.Println(len(grades))
	fruits := [...]int{1, 2, 3, 4, 5}
	fmt.Println(fruits)
	fmt.Println(len(fruits))
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}
	for index, value := range grades {
		fmt.Println(index, "=>", value)
	}
}
