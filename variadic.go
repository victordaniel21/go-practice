package main

import "fmt"

func numSum(num ...int) int {
	var number int
	for _, nums := range num {
		number += nums
	}
	return number

}

func numSub(num ...int) int {
	var number int
	for _, nums := range num {
		number -= nums
	}
	return number
}

func main() {
	totalSum := numSum(12, 15, 18)
	totalSub := numSub(100, 20, 30)
	fmt.Println(totalSum, totalSub)

}
