package main

import "fmt"

func main() {
	// as i know that slices are same as a array but the difference is that slices are dynamic in nature but arrays are static
	// so for creating slice don't give the size
	slice := []int{1, 2, 3}

	//see slice has capacity if you see there there is you didn't mentioned capacity so it is now 3 length of your slice
	fmt.Println("Capacity: ", cap(slice))
	n := 1
	fmt.Println("length is :", len(slice))
	for i := range slice {
		fmt.Println(n, "elements: ", slice[i])
		n++
	}

	//now append you want to append the element in your slice
	slice = append(slice, 4)
	slice = append(slice, 5)
	fmt.Println(slice)
	fmt.Println(cap(slice)) //now capacity become 6

	//you want to mention capacity then use make
	slice2 := make([]int, 0, 4)
	fmt.Println("length by using make ", len(slice2))
	fmt.Println("capacity by using make ", cap(slice2))
	fmt.Println("slice is: ", slice2)

	//internaly it has pointer that 1. points to array 2. Length 3. Capacity

}
