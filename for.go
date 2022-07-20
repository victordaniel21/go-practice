package main

import "fmt"

func main() {
	var x = 1
	for y := 1; y < 30; y++ {
		x = y * 2
		fmt.Println(x)
	}

}
