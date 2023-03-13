package main

import "fmt"

func main() {
	//nil
	//nil is similar to null in the other language
	//most golang type has default value, not nil. e.g bool(false), string(""), number<int/uint/float>(0) etc
	//nil is used on "non-primitive" data type (interface, map, slice, array, function, channel, pointer, etc.)

	var person map[string]any
	fmt.Println(person)
	fmt.Println(person == nil) //to check if map is nil
}
