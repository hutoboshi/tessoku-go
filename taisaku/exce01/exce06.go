package main

import "fmt"

func main() {
	a := 1071
	b := 1029
	ans := -1

	for {
		amari := a % b
		if amari == 0 {
			ans = b
			break
		}
		a = b
		b = amari
	}

	fmt.Println(ans)
}
