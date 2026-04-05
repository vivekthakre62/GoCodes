//generics are used for type safety and prevent the problem of the code duplication

package main

import "fmt"

func iterat[T any](items []T) {
	for _, item := range items {
		fmt.Println("item ", item)
	}
}

//you want some specific data type so

type stack2[T int | string | bool] struct {
	element []T
}

//comparable we can also use
type stack3[T comparable] struct {
	element []T
}

//struct
type stack[T any] struct {
	element []T
}

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []string{"Golang", "Java"}
	iterat(slice1)
	iterat(slice2)
	slice3 := []int{23, 3, 4, 5}
	myStack := stack[int]{
		element: slice3,
	}
	fmt.Println(myStack)
}
