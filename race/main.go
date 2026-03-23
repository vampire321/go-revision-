//
package main

import(
	"fmt"
	"sync"
)
//data race on counter 
//when you counter++ the cpu actually performs 3 steps Read(go and get the v)-Modify(update 50 to 51)-Write(copy new value back into memory) eventually what happens is since we are spawning 1,000 workers, two or more workers try to do work at the same time 
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
//to do := Once i must check for the questions of yesterday 