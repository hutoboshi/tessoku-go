package main

import "fmt"

func substrings(s string) []string {
	n := len(s)
	var substrs []string
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			substrs = append(substrs, s[i:j])
		}
	}
	return substrs
}

func main() {
	var s string
	fmt.Scan(&s)

	substrs := substrings(s)
	ansMap := make(map[string]bool)

	for _, substr := range substrs {
		ansMap[substr] = true
	}

	result := len(ansMap)
	fmt.Println(result)
}
