package main

import "fmt"

func main() {
	var a, b, w int
	fmt.Scan(&a, &b, &w)

	w_g := w * 1000

	minAns := 1_000_000_001  //十分大きな値を初期化とする
	maxAns := -1_000_000_001 //十分小さな値を初期値とする

	for x := 1; x <= 1_000_000; x++ {
		if a*x <= w_g && w_g <= b*x {
			minAns = min(minAns, x)
			maxAns = max(maxAns, x)
		}
	}

	if minAns == 1_000_000_001 {
		fmt.Println("UNSATISFIABLE")
	} else {
		fmt.Println(minAns, maxAns)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
