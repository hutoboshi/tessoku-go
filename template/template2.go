package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {

}

func inputArrInt() []int {
	scanner.Scan()
	inputs := strings.Split(scanner.Text(), " ")

	var ps []int
	for _, input := range inputs {
		p, _ := strconv.Atoi(input)
		ps = append(ps, p)
	}

	return ps
}

func inputArrString() []string {
	scanner.Scan()
	input := scanner.Text()
	inputs := strings.Fields(input)

	return inputs
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
