// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main(){
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        counts[input.Text()]++
    }
}

