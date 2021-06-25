package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	var input = "Hello, OTUS!"
	var output = stringutil.Reverse(input)
	fmt.Println(output)
}
