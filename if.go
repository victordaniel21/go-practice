package main

import "fmt"

func main() {
	var name string
	name = "Wayne"

	if name == "Daniel" {
		fmt.Println("welcome, Daniel!")
	} else if name == "Wayne" {
		fmt.Printf("Welcome %v\n", name)
	} else {
		fmt.Println("You are not in the list, please register first")
	}
}
