package main

import "fmt"

// this func accept x and y as parameter, and returns integer
func addTwoNumbers(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(addTwoNumbers(2, 5))
}
