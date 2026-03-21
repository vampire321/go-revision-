package main

import "fmt"

type Speaker interface {
	Speak() string 
}

type Dog struct{}

func (d Dog) Speak() string {
	return "woof!"
}

type Cat struct{}

func (c *Cat) Speak() string{
	return "Meow!"
}

func main() {
	var s Speaker
	d := Dog{}
	s=d
	fmt.Println(s.Speak())

	c := Cat{}
	s=&c
	fmt.Println(s.Speak())

}
