package main

import "fmt"

func main(){
	//buffered channel holds up to 3 values without a reciever
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3
	//ch <- 4 this would block because the buffer is full

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch) 

}