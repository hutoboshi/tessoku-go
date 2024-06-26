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

var buff = make([]byte, 100000)
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	// buff := make([]byte, 100000)
	// scanner.Buffer(buff, 100000)
	// input1 := inputArrInt()
	// n := int(input1[0])
	// arr := inputArrInt()
	// arrH := make([]int64, n)
	// for i := 0; i < n; i++ {
	// 	arrH[i] = int64(arr[i])
	// }

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

func inputArrInt() []int {
	scanner.Buffer(buff, 100000)
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
	scanner.Buffer(buff, 100000)
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
