package main

import (
	"fmt"
	"slices"
)

func main() {
	//creating a slice
	var s []int
	fmt.Println(s != nil)

	//mostly used way to create a slice
	var s1 = make([]int, 0, 5)
	fmt.Println(s1 == nil)
	//append element to slice
	s1 = append(s1, 10)
	s1 = append(s1, 20)
	s1 = append(s1, 30)
	s1 = append(s1, 40)
	s1 = append(s1, 50)
	fmt.Println(s1)

	//copying a slice
	var nums2 = make([]int, 0, 5)
	nums2 = append(nums2, 3)
	var nums3 = make([]int, len(nums2))
	copy(nums3, nums2)
	fmt.Println(nums2, nums3)

	//slice operator
	fmt.Println(s1[0:3])
	fmt.Println(s1[2:])

	//slices Equal mehtod
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	fmt.Println(slices.Equal(slice1, slice2))

	//2D slice
	nums4 := [][]int{{1, 2, 3}, {4, 6, 5}}
	fmt.Println(nums4)

}
