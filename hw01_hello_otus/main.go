package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	phrase := "Hello, OTUS!"
    reversed := reverse.String(phrase)
    fmt.Println(reversed)
}
