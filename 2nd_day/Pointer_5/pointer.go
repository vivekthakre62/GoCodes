package main

import "fmt"

//without pointer
func pointers(a int) {
	a = 5
	fmt.Println("Inside change", a)
}

//with pointer
func PointersWithPointer(a *int) {
	*a = 5
	fmt.Println("Inside change", *a)
}
func main() {
	a := 10
	pointers(a)

	fmt.Println("After Change pass by value", a)
	PointersWithPointer(&a)
	fmt.Println("After Change by using pass by reference", a)
}
