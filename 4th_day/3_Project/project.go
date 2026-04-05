package main

import (
	"fmt"
	"sync"
	"time"
)

type url struct {
	urls []string
}

func (u *url) getUrls(wg *sync.WaitGroup) {
	for _, url := range u.urls {
		wg.Add(1)

		go func(link string) {
			defer wg.Done()
			fmt.Println("Fetching:", link)
			time.Sleep(time.Second * 3)
		}(url)
	}
}
func main() {

	u := url{
		urls: []string{
			"https://localhost:8080/vivek.com",
			"https://localhost:8081/Rahul.com",
			"https://localhost:2300/Rajhansh.com",
			"https://localhost:7788/mynews.com",
			"https://localhost:9090/anue.com",
		},
	}
	var wg sync.WaitGroup
	u.getUrls(&wg)
	wg.Wait()
	fmt.Println("All URLs Done")
}
