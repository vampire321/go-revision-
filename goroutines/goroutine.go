package main

import(
	"fmt"
	"sync"
	"time"
)

func downloadfile(filename string, wg *sync.WaitGroup){
	defer wg.Done() //withput these the program hangs for ever at wg.Wait() because the cpunter never reaches the zero (this is called the deadlock)
	fmt.Println("started downloading", filename)
	time.Sleep(300 * time.Millisecond)
	fmt.Println("finished downloading", filename)
}

func main(){
	var wg sync.WaitGroup
	files := []string{
		"file.txt",
		"image.png",
		"data.csv",
		"video.mpo4",
	}
	for _,file := range files{
		wg.Add(1) //this is used to increment the counter of the waitgroup, so that the main() can wait for all the goroutines to finish before exiting
		go downloadfile(file, &wg)
	}
	fmt.Println("All downloads are started: waiting")
	wg.Wait() //without this line , the main() exists instantly, killing all the goroutines
	fmt.Println("all downloads complete")
}
	