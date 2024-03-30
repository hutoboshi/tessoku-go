package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Split(bufio.ScanWords)

	var n int
	n = nextInt()

	mountain := make(map[int]string)
	keys := make([]int, n)
	for i := 0; i < n; i++ {
		var tmp1 string
		var tmp2 int
		tmp1, tmp2 = nextString(), nextInt()
		mountain[tmp2] = tmp1
		keys[i] = tmp2
	}

	sort.Ints(keys)

	fmt.Println(mountain[keys[len(keys)-2]])
}

func nextInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func nextString() string {
	scanner.Scan()
	test := scanner.Text()
	return test
}
