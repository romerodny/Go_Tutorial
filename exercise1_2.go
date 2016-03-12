// Exercise 1.2 of the Go Programming Language book
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sep, s string
	for i := 1; i < len(os.Args); i++ {
		temp := strconv.Itoa(i)
		s += sep + temp + " " + os.Args[i]
		sep = "\n"
	}
	
	fmt.Println(s)
}