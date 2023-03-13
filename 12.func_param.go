package main

import "fmt"

// this func accept x and y as parameter, and returns integer
func addTwoNumber(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(addTwoNumber(2, 5))
}
