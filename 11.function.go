package main

import (
	"fmt"
	"strings"
)

func printHelloWorldMultipleTime(multipler int) {
	fmt.Println(strings.Repeat("Hello World\n", multipler))
}

func main() {
	printHelloWorldMultipleTime(10)
}
