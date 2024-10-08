package main

import "fmt"

func main() {
	num := 387420489.0
	for i := 0; i < 48; i++ {
		num = num * 0.8
	}

	fmt.Println(num)
}
