package main

import (
	"fmt"
	"strconv"
)

func main() {
	count := 0
	for i := 1; i <= 999999999; i++ {
		if i%4818 == 0 && isMonotonic(i) {
			count++
		}
	}

	fmt.Println(count)
}

func isMonotonic(n int) bool {
	s := strconv.Itoa(n)
	incre := true
	decre := true
	for i := 1; i < len(s); i++ {
		if s[i] > s[i-1] {
			decre = false
		}
		if s[i] < s[i-1] {
			incre = false
		}
	}

	return incre || decre
}
