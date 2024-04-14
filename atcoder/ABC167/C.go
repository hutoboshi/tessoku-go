package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	input1 := inputArrInt()
	n := input1[0]
	m := input1[1]
	x := input1[2]

	arrC := make([]int, n)
	arrA := make([][]int, n)
	for i := 0; i < n; i++ {
		input2 := inputArrInt()
		arrC[i] = input2[0]
		arrA[i] = input2[1:]
	}

	ans := math.MaxInt32
	for msk := 0; msk < (1 << n); msk++ {
		cost := 0
		understanding := make([]int, m)

		for i := range arrC {
			if msk&(1<<i) != 0 {
				cost += arrC[i]
				for j := range arrA[i] {
					understanding[j] += arrA[i][j]
				}
			}
		}

		ok := true
		for _, val := range understanding {
			if val < x {
				ok = false
				break
			}
		}

		if ok {
			ans = min(ans, cost)
		}
	}

	if ans == math.MaxInt32 {
		ans = -1
	}

	fmt.Println(ans)
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
