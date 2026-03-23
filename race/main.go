package main

import(
	"fmt"
	"sync"
)
//data race on counter 
var ( 
	counter int
	mu sync.Mutex
)

func increment(wg *sync.WaitGroup){
	defer wg.Done()
	mu.Lock() // accquire the lock - other goroutines block here
	counter++ //safe -only one goroutine runs this at a time
	mu.Unlock() //release the lock
}

func main(){
	var wg sync.WaitGroup
	for i := 0 ; i< 1000 ; i++{
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("counter:",counter) //it should print 1000, often isnt
}