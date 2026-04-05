package main

import "fmt"

//closure is a function that can access values outside of its scope and it can also modify those values.

func counter() func() int {
	var count int = 0
	return func() int {
		count += 1
		return count
	}
}
func main() {
	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())
}
