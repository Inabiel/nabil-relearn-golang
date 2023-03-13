package main

import "fmt"

func incrementer() func() int {
	counter := 0
	return func() int {
		counter += 1
		return counter
	}
}

func main() {
	count := incrementer()
	fmt.Println(count())
}
