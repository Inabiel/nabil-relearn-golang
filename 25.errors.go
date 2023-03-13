package main

import (
	"errors"
	"fmt"
)

func divider(number, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("Cannot divide number by zero")
	} else {
		return number / divider, nil
	}
}

func main() {
	hasil, err := divider(10, 2)

	if err == nil {
		fmt.Println("hasil", hasil)
	} else {
		fmt.Println("err", err.Error())
	}
}
