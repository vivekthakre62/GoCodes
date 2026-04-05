package main

import (
	"fmt"
	"sync"
)

type post struct {
	view int
	//for prevent the problem of Race condition means multiple thread use single resource at a time then inconsistent data will be calculate so we use mutex
	mu sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		p.mu.Unlock()
	}()

	p.mu.Lock()
	p.view += 1
	//p.mu.Unlock() *** so if any error will be come in the prev line then the unlock function will not call and goroutine stil is in the locked situation so for this we can do one thing we have to use this function inside the defer

}

func main() {
	p := post{
		view: 0,
	}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go p.inc(&wg)
	}
	wg.Wait()
	fmt.Print(p.view)
}
