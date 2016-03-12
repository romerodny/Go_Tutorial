// Exercise 1.3 from the Go Programming Language book
package main

import (
    "fmt"
    "os"
    "testing"
    "strings"
)

func main() {
    version1()
    version2()
    version3() 
}

func version1() {
    var s, sep string
    for i:= 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Println(s)
}

func version2(){
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}

func version3() {
    fmt.Println(strings.Join(os.Args[1:], " "))
}

func Benchmarkversion1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        version1()
    }    
}

