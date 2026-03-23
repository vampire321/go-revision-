package main

import(
	"fmt"
	"sync"
)
//data race on counter 
var counter int

func increment(wg *sync.WaitGroup){
	defer wg.Done()
	counter++
}

func main(){
	var wg sync.WaitGroup
	for i := 0;i< 1000 ; i++{
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("counter:",counter) //it should print 1000, often isnt
}