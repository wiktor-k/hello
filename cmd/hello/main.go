// Prints sample string serving as a demonstration that
// go framework works as expected.
// See: https://en.wikipedia.org/wiki/%22Hello,_World!%22_program
package main

import (
    "fmt"
    "metacode.biz/hello"
)

func main() {
    fmt.Printf("hello, world\n")
    fmt.Print(hello.Double(5011))
}
