package main

import "fmt"

const pi = 3.14

func main() {
	const gravity = 9.8
	fmt.Println(pi)
	fmt.Println(gravity)

	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port)
}
