package main

import "fmt"

func main() {
	// grades := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(grades)
	// fmt.Println(len(grades))
	// fruits := [...]int{1, 2, 3, 4, 5}
	// fmt.Println(fruits)
	// fmt.Println(len(fruits))
	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Println(fruits[i])
	// }
	// for index, value := range grades {
	// 	fmt.Println(index, "=>", value)
	// }
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := arr[1:8]
	fmt.Println(slice)
	sub_slice := slice[0:3]
	fmt.Println(sub_slice)

	slice1 := make([]int, 5, 8)
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

}
