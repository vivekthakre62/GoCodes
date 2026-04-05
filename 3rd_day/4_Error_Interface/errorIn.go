package main

import (
	"errors"
	"fmt"
)

// we can create own custom error so
type MyError struct {
	message string
}

func (myErr MyError) Error() string {
	return myErr.message
}

func divide(a, b float64) (float64, error) {
	ans := a / b
	if b == 0 {
		return 0, errors.New("Can't divide by 0") //here New is function that take string and convert it to an errors.errorStrin and return as an error val
	}
	return ans, nil
}

func random(s string) (string, MyError) {
	if s == "Y" {
		return "", MyError{
			message: "Error!, Danger",
		}
	} else {
		return "No Error!", MyError{
			message: "Everything is Good!",
		}
	}
}
func main() {
	//error interface is builtin interface in java
	ans, err := divide(10, 2) //that will give error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Division is: ", ans)
	}

	ans1, err1 := random("N")
	fmt.Println(ans1, err1)

	//fmt also have Errorf method for showing error it return a string
	f := -1
	if f <= 0 {
		err := fmt.Errorf("Error ! %d", f)
		fmt.Println(err)
	}

}
