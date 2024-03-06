package main

import (
	"fmt"
	"math"
)

func main() {
	var q int
	fmt.Scan(&q)

	arrX := make([]int, q+1)
	for i := 1; i <= q; i++ {
		fmt.Scan(&arrX[i])
	}

	for i := 1; i <= q; i++ {
		ans := isPrime(arrX[i])
		if ans {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func isPrime(x int) bool {
	if x <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
