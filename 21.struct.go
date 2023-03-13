package main

import "fmt"

type People struct {
	FirstName, LastName string
	Age                 int
	Gender              bool
}

// struct Method
// way to "inject" function into struct
func (people *People) combineName() string {
	return people.FirstName + " " + people.LastName
}

// there are two ways to define the type, either using the Type itself or to use the pointer to type.
// if you use the pointer(*), you can mutate the struct.
// if not, you're only mutating the struct in method only
// this is because golang is passed by value, not reference
func (people *People) changeFirstName(changedName string) {
	people.FirstName = changedName
}

func main() {
	//struct
	// is a data template to combine one or more data
	//in usual programming term, this would be closest to object in a OOP manners (js, java, php, etc.)

	firstMf := People{FirstName: "Nabiel", LastName: "Izzullah Pansuri", Age: 21, Gender: true}
	if firstMf.Gender == true {
		fmt.Println(firstMf.combineName(), "gender is male")
	} else {
		fmt.Println(firstMf.combineName(), "gender is female")
	}
	firstMf.changeFirstName("Jonathan")
	fmt.Println(firstMf.combineName())
}
