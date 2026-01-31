package main

import "fmt"

func greet(target string) {
    fmt.Printf("Hello, %s!\n", target)
}

func main() {
    greet("Literate Programmer")
}