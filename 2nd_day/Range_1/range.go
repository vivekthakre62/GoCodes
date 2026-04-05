package main

import "fmt"

func main() {
	//range
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for i, num := range nums {
		fmt.Println(i, num)
		sum += num

	}
	fmt.Println(sum)

	//range on map
	maps := map[string]int{"Aman": 10, "Raman": 20, "Suman": 30}
	maps["Pawan"] = 40
	for k, v := range maps {
		fmt.Println(k, "=", v)
	}
	for k, v := range map[string]int{"Asus": 30000, "Dell": 40000, "HP": 50000} {
		fmt.Println("Price of", k, " is ", v)
	}
	//string are also iterable
	for k, v := range "Hello" {
		fmt.Println(k, " = ", v)
		//ascii to character
		fmt.Println(k, " = ", string(v))

	}

}
