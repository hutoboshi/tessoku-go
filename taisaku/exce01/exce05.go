package main

import "fmt"

func main() {
	count := 0
	for i := 1; i <= 1000; i++ {
		if romanLen(i) == 9 {
			count++
		}
	}

	fmt.Println(count)
}

func romanLen(n int) int {
	numerals := map[int]string{
		1000: "M", 900: "CM", 500: "D", 400: "CD", 100: "C", 90: "XC", 50: "L", 40: "XL",
		10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I",
	}

	result := ""
	for value, numeral := range numerals {
		for n >= value {
			result += numeral
			n -= value
		}
	}
	// if len(result) == 9 {
	// 	fmt.Println(result)
	// }
	return len(result)
}
