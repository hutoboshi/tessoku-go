package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const INF = 1 << 60

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	// input1 := inputArrInt()
	// n := int(input1[0])
	// arr := inputArrInt()
	// arrH := arr
	var n int
	fmt.Scan(&n) // 入力：N

	arrH := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arrH[i]) // 入力：h[i]
	}

	dp := make([]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = INF
	}

	dp[0] = 0

	//ループ:DPテーブルを計算
	for i := 1; i < n; i++ {
		chmin(&dp[i], dp[i-1]+int64(math.Abs(float64(arrH[i]-arrH[i-1]))))
		if i > 1 {
			chmin(&dp[i], dp[i-2]+int64(math.Abs(float64(arrH[i]-arrH[i-2]))))
		}
	}

	fmt.Println(dp[n-1])
}

// chmaxはaをbで更新できるなら更新し、その結果を返す関数
func chmax(a *int64, b int64) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}

// chmin はaをbでこうしんできるなら更新し、その結果をを返す関数
func chmin(a *int64, b int64) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

func inputArrInt() []int64 {
	scanner.Scan()
	input := scanner.Text()
	inputs := strings.Fields(input)

	var ps []int64
	for _, input2 := range inputs {
		p, _ := strconv.ParseInt(input2, 10, 64)
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
