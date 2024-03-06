package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	arrT := make([]string, n+1)
	arrA := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		var t string
		fmt.Scan(&t, &arrA[i])
		arrT[i] = t
	}

	var ans int64 = 0
	for i := 1; i <= n; i++ {
		switch arrT[i] {
		case "+":
			ans += arrA[i]
		case "-":
			ans -= arrA[i]
		case "*":
			ans *= arrA[i]
		}

		if ans < 0 {
			ans += 10000
		}

		ans %= 10000
		fmt.Println(ans)
	}
}
