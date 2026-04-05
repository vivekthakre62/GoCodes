package main

import "fmt"

func add(a, b int) int {
	return a + b
}

//you can also return multiple values from a function

func addAndSubtract(a, b int) (int, int) {
	return a + b, a - b
}

// we can pass function as an argument to another function

func myValue() func(a int) int {
	return func(a int) int {
		return 4
	}
}
func MyValue2(fn func(a int) int) int {
	return fn(10)
}
func main() {
	values := func(a int) int {
		return a * 2
	}
	fmt.Println(MyValue2(values))
	fmt.Println(myValue()(10))
	ans := add(10, 20)
	ans1, ans2 := addAndSubtract(20, 10)
	fmt.Println(ans)
	fmt.Println(ans1, ans2)
}
