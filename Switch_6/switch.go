package main

import (
	"fmt"
	"time"
)

func main() {
	// day := "Monday"

	// switch day {
	// case "Monday":
	// 	fmt.Println("Monday")

	// case "Tuesday":
	// 	fmt.Println("Tuesday")
	// case "Wednesday":
	// 	fmt.Println("Wednesday")
	// case "Thursday":
	// 	fmt.Println("Thursday")
	// case "Friday":
	// 	fmt.Println("Friday")
	// case "Saturday":
	// 	fmt.Println("Saturday")
	// case "Sunday":
	// 	fmt.Println("Sunday")
	// default:
	// 	fmt.Println("Other")
	// }

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("WorkDay")
	}

	//type switch
	myType := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("Integer")
		case string:
			fmt.Println("String")
		case bool:
			fmt.Println("Boolean")
		default:
			fmt.Println("Other")
		}
	}
	myType(true)
}
