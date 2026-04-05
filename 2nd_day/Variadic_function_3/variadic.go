package main

import "fmt"

//variadic function is a function that can take variable number of arguments

//variadic function

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
func main() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(sum(nums...))
	fmt.Println(sum(1, 2, 3, 4, 5))
}
