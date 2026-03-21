package main

import (
    "flag"
    "fmt"
)

func main() {
    // flag.Int returns *int (a pointer to int), not int
    // arguments: flag name, default value, description
    top    := flag.Int("top", 10, "show top N words")
    upper  := flag.Bool("upper", false, "print words in uppercase")
    prefix := flag.String("prefix", "", "only count words starting with this")

    // MUST call flag.Parse() before reading any flag values
    // This reads os.Args and fills in the values
	
    flag.Parse()

    // Dereference the pointer with * to get the actual value
	fmt.Println("top:", *top)
    fmt.Println("upper:", *upper)
    fmt.Println("prefix:", *prefix)

    // flag.Args() returns everything AFTER the flags
    fmt.Println("remaining args:", flag.Args())
}