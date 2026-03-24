//Sender and reciver go routines
package main

import "fmt"

func producer(ch chan<- string){
	// chan<- means "send-only channel" — can only send, not receive
	urls := []string{"google.com", "github.com", "stripe.com"}
	for _, url := range urls {
        fmt.Println("sending:", url)
		ch <- url //sernding url into the channel
	}
	close(ch) //signals:=no more values will be sent
}

func consumer(ch <-chan string) {
	// <-chan means "receive-only channel" - only can recieve
	for urls := range ch {
		//Recives untill the channel is closed 
		fmt.Println("recieved:", urls)
	}

	fmt.Println("channel closed,consumer done")
}

func main() {
	//made an unbufferd channel of strings and converted into buffered channel
	ch := make(chan string , 3)

	go producer(ch) //without go it blocks waiting for a reciver because both producer and consumer are at main , but there is no goroutine to recieve 
	consumer(ch)
}