package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Create a buffered channel with capacity 3
	ch := make(chan int, 3)

	// 2. Spawn the Producer
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Printf("Producer: Sending %d to the channel...\n", i)
			ch <- i
			// Notice how fast this prints until the buffer (3) is full!
		}
		// 3. Don't forget to close!
		close(ch)
	}()

	// 4. Main is the "Slow Consumer"
	for num := range ch {
		fmt.Printf("Main: Received %d (Processing...)\n", num)
		
		// Simulate a slow DB write (500ms)
		time.Sleep(500 * time.Millisecond) 
	}

	fmt.Println("All metrics processed!")
}