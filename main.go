package main

import (
	"fmt"

	"github.com/rotemtam/protoc-gen-go-ascii/example"
)

func main() {
	ex := &example.Hello{}
	fmt.Println("going to print example.Hello's ASCII art representation:")
	fmt.Println(ex.Ascii())
}