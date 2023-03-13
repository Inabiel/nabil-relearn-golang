package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println(counter)
		counter += 1
	}

	obj := map[string]any{
		"first":  1,
		"second": 2,
		"third":  3,
	}

	for k, v := range obj {
		fmt.Println("key is", k, "with value", v)
	}

	//for with short statement
	//it has var declaration, init statement, and post statement
	for secondCounter := 1; secondCounter <= 11; secondCounter += 1 {
		fmt.Println(secondCounter)
	}

	arr := [5]int{1, 2, 3, 4, 5}

	//old school for style
	for i := 0; i < len(arr); i += 1 {
		fmt.Println(arr[i])
	}

	//we can use underscore to remove unwanted item in looping
	for _, v := range obj {
		fmt.Println("value is", v)
	}
}
