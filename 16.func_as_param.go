package main

import (
	"fmt"
	"strings"
)

func sayHelloDependingOntime(timeOfDay string, filter func(string) string) string {
	return ("Good " + filter(timeOfDay) + "!")
}

func theTimeOfDay(time string) string {
	switch strings.ToLower(time) {
	case "morning":
		return "Morning"
	case "afternoon":
		return "Afternoon"
	case "evening":
		return "Evening"
	default:
		return "Night"
	}
}

func main() {
	//function can be supplied into another func! thats crazy!

	whatTimeIsThis := sayHelloDependingOntime("morning", theTimeOfDay)
	fmt.Println(whatTimeIsThis)

}
