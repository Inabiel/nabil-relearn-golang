package main

import "fmt"

func sumAllFromSlice(nums ...int) int {
	total := 0
	for _, i := range nums {
		total += i
	}
	return total
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(sumAllFromSlice(arr...))

}
