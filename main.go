package main
import (
	"flag"
	"fmt"
)
func main(){
	verbose := flag.Bool("verbose",false,"display logs")
	flag.Parse()
	if *verbose{
		fmt.Println("Verbose mode ON: showing logs...")
	} else {
		fmt.Println("Verbose mode OFF ")
	}
	fmt.Println("value of port:", *nFlag)
	fmt.Println("user:", *user)
	variable()
	structs()
}