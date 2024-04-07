package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	arrX := make([]int, n)
	arrY := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arrX[i], &arrY[i])
	}

	for i := 0; i < n; i++ {
		var ans int
		var maxnum float64 = 0
		for j := 0; j < n; j++ {
			num := euclideanDistance(float64(arrX[i]), float64(arrY[i]), float64(arrX[j]), float64(arrY[j]))
			if maxnum < num {
				maxnum = num
				ans = j + 1
			}
		}
		fmt.Println(ans)
	}
}

func euclideanDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
