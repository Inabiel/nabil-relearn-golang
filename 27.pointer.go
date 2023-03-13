package main

import "fmt"

type Peoples struct {
	FirstName, LastName string
	Age                 int
	Gender              bool
}

func main() {
	//golang uses pass-by-value, rather than pass-by-reference such as JS, python, and other dynamically typed lang
	//if we pass the param into func, it will copy the value into memory address, which makes it hard
	//it will become hard to trace too
	//the solution is, we need to mutate it using pointer, it will mutate it in the mem addr
	//example below

	address1 := Peoples{FirstName: "Nabiel", LastName: "Izzullah", Age: 21, Gender: true}
	address2 := &address1
	*address2 = Peoples{"Nabil", "2", 23, false}
	fmt.Println(address1)  // address1 doesnt change! strange
	fmt.Println(*address2) //now address1 will change too
}
