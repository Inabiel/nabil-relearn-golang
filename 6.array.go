package main

import "fmt"

func main() {
	//Array is a type data consists of values with the same type
	//Array cannot be mixed type
	//Array size is fixed, similar to c-style array
	//Need to allocate the size before inserting values
	//to initialize array, use [<lengthOfArray>]<typeOfArray>
	//to insert it, we can do <var>[<index>] = <value>

	var someArray [4]int64
	someArray[0] = 2131231
	someArray[2] = 321312
	fmt.Println(someArray)

	var newArray = [5]string{
		"this", "is", "an", "array",
	}

	fmt.Println(newArray)

}
