package main

import "fmt"

func main() {
	var nums [5]int
	fmt.Println(len(nums))
	//add new value
	nums[0] = 10
	//getAll value
	fmt.Println(nums)
	//getASingleValue
	fmt.Println(nums[0])
	//boolean array
	var bools [3]bool
	bools[0] = true
	fmt.Println(bools)

	names := [3]string{"vivek", "Sachin", "Mahi"}
	fmt.Println(names)

	//2D array
	matrix := [2][2]int{{1, 2}, {1, 2}}
	fmt.Println(matrix)

}
