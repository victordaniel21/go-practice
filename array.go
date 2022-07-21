package main

import "fmt"

func main() {
	//Find the midian
	x := [7]float32{23, 24, 30, 32, 27, 30, 25} //data collected
	var Midian float32
	for i := 0; i < 7; i++ {
		Midian += x[i]

	}
	fmt.Println(Midian / float32(len(x)))
}
