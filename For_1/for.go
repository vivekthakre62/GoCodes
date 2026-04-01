package main

import "fmt"

func main() {

	//while loop implementation in go
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//infinite loop
	// for {
	// 	println(1)
	// }

	//tradition for loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//range loop
	for i := range 11 {
		fmt.Println(i)
	}

}
