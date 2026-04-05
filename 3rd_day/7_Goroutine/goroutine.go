package main

import (
	"fmt"
	"sync"
)

func count(i int, w *sync.WaitGroup) {

	defer w.Done() //it decrement waitGroup task by 1.
	fmt.Println("Doing a task ", i)
}

//	func count(i int) {
//		fmt.Println("doing a task ", i)
//	}
func main() {
	// for i := 1; i <= 10; i++ {
	// 	go func() {
	// 		fmt.Println(i)
	// 	}()

	// }
	// time.Sleep(2 * time.Second)

	var weg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		weg.Add(1) //Add adds delta, which may be negative, to the waitgroup task counter.
		count(i, &weg)
	}
	weg.Wait() //wait until waitGroupTask Become 0
}
