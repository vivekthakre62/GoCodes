package main

import (
	"fmt"
	"time"
	// "math/rand"
	// "time"
)

func sendEmail(emailChan chan string, done chan bool) {
	defer func() { done <- true }()

	for email := range emailChan {
		fmt.Println("Email Send To", email)
		time.Sleep(time.Second)
	}

}
func processing(res chan bool) {
	defer func() { res <- true }()
	fmt.Println("Proceesing....")
}
func add(result chan int, a, b int) int {

	newResult := a + b
	result <- newResult
	return <-result
}
func process(newChannel chan int) {
	for p := range newChannel {
		fmt.Println("Process numbers ", p)

	}
}

func main() {

	// newChannel := make(chan int)
	// result := make(chan int)

	// go process(newChannel)
	// for {
	// 	newChannel <- rand.Intn(100)
	// 	time.Sleep(time.Second * 2)
	// }

	// go add(result, 4, 5)
	// res := <-result  // when we use channels then it is blocking the sending and recieving
	// fmt.Println(res)

	// res := make(chan bool)
	// go processing(res)
	// <-res
	// emailChan := make(chan string, 100)
	// done := make(chan bool)
	// go sendEmail(emailChan, done)
	// for i := 1; i < 5; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// 	// time.Sleep(time.Second)
	// }
	// //we need to close buffered channel
	// close(emailChan)

	// fmt.Println("Done Sending")
	// <-done

	//work on multiple channels
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()
	go func() {
		chan2 <- "pong"
	}()

	//for getting the data we have to use for loop and select
	for i := 0; i < 2; i++ {
		//select
		select {
		case chanVal1 := <-chan1:
			fmt.Println("value of channel 1 is:", chanVal1)
		case chanVal2 := <-chan2:
			fmt.Println("value of channel 2 is:", chanVal2)
		}

	}

}
