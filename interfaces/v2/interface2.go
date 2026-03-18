package main

import (
    "fmt"
    "strings"
)

// Formatter is an interface
// Any type with a Format(string) string method satisfies it
// No "implements" keyword needed — it's automatic
type Formatter interface {
    Format(word string) string
}

// UpperFormatter satisfies Formatter
type UpperFormatter struct{}
func (f UpperFormatter) Format(word string) string {
    return strings.ToUpper(word)
}

// PrefixFormatter also satisfes formatter
type PrefixFormatter struct {
    Prefix string
}

func (f PrefixFormatter) Format(word string) string {
    return f.Prefix + word
}

// NoopFormatter — does nothing, returns word as-is
type NoopFormatter struct{}
func (f NoopFormatter) Format(word string) string {
    return word
}

// printWords accepts ANY type that satisfies Formatter
// It doesn't know or care if it's Upper, Prefix, or Noop
func printWords(words []string, f Formatter) {
    for _, w := range words {
        fmt.Println(f.Format(w))
    }
}

func main() {
    words := []string{"hello", "world", "golang"}

    fmt.Println("--- uppercase ---")
    printWords(words, UpperFormatter{})

    fmt.Println("--- prefixed ---")
    printWords(words, PrefixFormatter{Prefix: ">> "})

    fmt.Println("--- unchanged ---")
    printWords(words, NoopFormatter{})
}