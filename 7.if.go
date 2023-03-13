package main

import "fmt"

func main() {
	var someArray = [5]bool{
		true, false, false, true, false,
	}

	if len(someArray) > 3 {
		fmt.Println("array is longer than 3")
	}

	//If with Short Statement
	//Short Statement means we can declare variable right before if statement, making it more concise
	if length := len(someArray); length > 3 {
		fmt.Println("array is longer than 3")
	}
}
