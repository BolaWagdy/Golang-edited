package main

import "fmt"

func sqrt(x float32) float32 {
	var z float32 = 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z

}

func main() {
	fmt.Println(sqrt(9))
	fmt.Println("hello world")

}
