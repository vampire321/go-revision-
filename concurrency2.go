package main

import(
	"fmt"
	"sync"
)

func checkUrl(URL string , wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("checking :",URL)
}

func main(){
	urls := []string{
		"https://google.com",
		"https://github.com",
	}

	var wg sync.WaitGroup

	for _,url:= range urls{
		wg.Add(1)
		go checkUrL(url , &wg)
	}

	wg.Wait()
	fmt.Println("all checks completed")
}