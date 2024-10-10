//フィボナッチ数列

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	binarySearch2(n)
}

func binarySearch2(n int) {
	left := 1
	right := 1000000

	cnt := 0

	for right-left+1 >= 1 {
		fmt.Println(left, right)
		mid := (left + right) / 2
		fmt.Println("mid", mid)
		midstr := strconv.Itoa(mid)
		doubleMid, _ := strconv.Atoi(midstr + midstr)
		if doubleMid == n {
			cnt = mid
			break
		}
		if doubleMid < n {
			left = mid + 1
			cnt = mid
			fmt.Println("一致", cnt)
		}
		if doubleMid > n {
			right = mid - 1
		}
	}

	fmt.Println(cnt)
}
