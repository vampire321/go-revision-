package main

import(
	"fmt"
)
func sending(ch1 chan<- int){
		nums := []int{1,2,3,4,5,6}
		for num := range nums{
			fmt.Println("sending:", num)
			ch1 <- num
		}
		close(ch1)
	}
	func main(){
	ch1 := make(chan int, 5)
	go sending(ch1)
	for number := range ch1{
		fmt.Println("received:", number)
	}
	fmt.Println("All numbers processed!")
}