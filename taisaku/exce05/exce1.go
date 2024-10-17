package main

import (
	"fmt"
	"strconv"
)

func main() {
	cnt := 0
	for i := 1; i < 999999999; i++ {
		if i%4818 == 0 && isMonotonic2(i) {
			cnt++
		}
	}

	fmt.Println(cnt)
}

func isMonotonic2(n int) bool {
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
