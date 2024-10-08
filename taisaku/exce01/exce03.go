package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	count := 0
	// for i := 1; i <= 3000; i++ {
	// 	if i%3 == 0 {
	// 		count++
	// 		continue
	// 	}
	// 	strI := strconv.Itoa(i)
	// 	for _, v := range strI {
	// 		if v == '3' {
	// 			count++
	// 			break
	// 		}
	// 	}
	// }

	for i := 1; i <= 3000; i++ {
		if i%3 == 0 || countainsDigtThree(i) {
			count++
		}
	}

	fmt.Println(count)
}

func countainsDigtThree(n int) bool {
	return strings.Contains(strconv.Itoa(n), "3")
}
