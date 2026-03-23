package main

import (
	"fmt"
	"sync"
)

func checkUrL(url string, wg *sync.WaitGroup) {
	defer wg.Done() // signals "I am finished" when this function returns
	// defer means: run this line when the function exits, no matter what
	fmt.Println("checking:", url)
}

func Concurrency() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://stripe.com",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)             // tell WaitGroup: one more goroutine starting
		go checkUrL(url, &wg) // launch goroutine — pass &wg (pointer)
	}

	wg.Wait() // BLOCK here until all goroutines call wg.Done()
	fmt.Println("all checks complete")
}
