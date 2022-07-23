package main

import "fmt"

func numbers(num ...int) int {
	var number int
	for _, nums := range num {
		number += nums
	}
	return number

}

func main() {
	total := numbers(12, 15, 18)
	fmt.Println(total)

}
