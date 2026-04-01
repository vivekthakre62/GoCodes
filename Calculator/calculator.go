package main

import "fmt"

func main() {
	var a, b int
	for {
		var inp string
		fmt.Println("Do you want to use calculator? (Y/N)")
		fmt.Scanln(&inp)
		if inp == "N" || inp == "n" {
			break
		} else {
			fmt.Println("Enter first number")
			fmt.Scanln(&a)
			fmt.Println("Enter Second Number")
			fmt.Scanln(&b)
			fmt.Println("Select Operation")
			fmt.Println("1.Addition")
			fmt.Println("2.Subtraction")
			fmt.Println("3.Multiplication")
			fmt.Println("4.Division")
			var choice int
			fmt.Scanln(&choice)
			switch choice {
			case 1:
				fmt.Println("Your Answer is: ", a+b)
			case 2:
				fmt.Println("Your Answer is: ", a-b)
			case 3:
				fmt.Println("Your Answer is: ", a*b)
			case 4:
				fmt.Println("Your Answer is: ", a/b)

			default:
				fmt.Println("Input is wrong")
			}
		}

	}
}
