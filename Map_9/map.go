package main

import (
	"fmt"
	"maps"
)

func main() {
	//creating a map
	ages := make(map[string]int)
	ages["Alice"] = 20
	ages["vivek"] = 22
	fmt.Println(ages["Alice"])

	delete(ages, "vivek")
	fmt.Println(ages["vivek"])
	m := map[string]int{"Alice": 20, "Vivek": 22}
	fmt.Println(m["Alice"])

	m1 := map[string]int{"Price": 40000, "Phone": 4}

	m2 := map[string]int{"Price": 40000, "Phone": 4}

	_, ok := m1["Price"]
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	fmt.Println(maps.Equal(m1, m2))
}
