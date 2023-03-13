package main

import "fmt"

func main() {

	var someArray = [10]string{
		"this", "is", "an", "array", "filled", "with", "values",
	}

	//Switch with short statement
	switch firstArray := someArray[0]; firstArray {
	case "this":
		fmt.Println("value is <this>")
	default:
		fmt.Println("value is unknown")
	}

	//Switch can be used as a short if too!
	length := len(someArray)
	switch {
	case length > 6:
		fmt.Println("array is long")
	default:
		fmt.Println("array is short")
	}

}
