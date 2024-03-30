package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Split(bufio.ScanWords)

	n, k := nextString(), nextInt()
	var ans string = n

	for i := 1; i <= k; i++ {
		ans = calcN(ans)
	}

	fmt.Println(strconv.Atoi(ans))
}

func nextInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func nextString() string {
	scanner.Scan()
	text := scanner.Text()
	return text
}

func calcN(sN string) string {
	arr := make([]int, len(sN))
	for i, v := range sN {
		arr[i], _ = strconv.Atoi(string(v))
	}
	sort.Ints(arr)

	var g1, g2 int

	for i := 0; i < len(arr); i++ {
		g1 += int(math.Pow(float64(10), float64(i))) * arr[i]
		g2 += int(math.Pow(float64(10), float64(i))) * arr[len(arr)-i-1]
	}

	summary := g1 - g2
	if summary <= 0 {
		summary = 0
	}

	return string(summary)
}
