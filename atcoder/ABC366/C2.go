package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	Q, _ := strconv.Atoi(scanner.Text())

	countMap := make(map[int]int)
	uniqueCount := 0

	for i := 0; i < Q; i++ {
		scanner.Scan()
		query := strings.Fields(scanner.Text())

		switch query[0] {
		case "1":
			x, _ := strconv.Atoi(query[1])
			if countMap[x] == 0 {
				uniqueCount++
			}
			countMap[x]++
		case "2":
			x, _ := strconv.Atoi(query[1])
			countMap[x]--
			if countMap[x] == 0 {
				uniqueCount--
			}
		case "3":
			fmt.Println(uniqueCount)
		}
	}
}
