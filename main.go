package main

import "fmt"

func checkURL(url string){
	fmt.Println("checking:",url)
}

func main(){
	urls := []string{
		"https://google.com",
		"https://google.com",
		"https://google.com",
	}

	for _,url := range urls {
		go checkURL(url) //launch each check currently
	}

	fmt.Println("all checks launched")
}