// Exercise 1.1 of the Go Programming Language book
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	method1()
	method2()
	method3()
	method4()
}

func method1() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func method2() {
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	
	fmt.Println(s)
}

func method3() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

func method4() {
	fmt.Println(os.Args[0:])
}


