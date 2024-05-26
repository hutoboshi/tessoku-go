package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	//入出力関連
	scan := bufio.NewReader(os.Stdin)
	print := bufio.NewWriter(os.Stdout)
	defer print.Flush()

	var n, q int

	fmt.Fscan(scan, &n, &q)

	boxes := make([][]int, 200020)

	cards := make([][]int, 200020)

	for q > 0 {
		q--
		var h int
		fmt.Fscan(scan, &h)

		if h == 1 {
			var j, k int

			fmt.Fscan(scan, &j, &k)

			boxes[k] = append(boxes[k], j)
			cards[j] = append(cards[j], k)

		} else if h == 2 {
			var k int
			fmt.Fscan(scan, &k)
			sort.Ints(boxes[k])

			for _, v := range boxes[k] {
				fmt.Fprintln(print, v)
			}
		} else {
			var k int
			fmt.Fscan(scan, &k)

			uniqCardNumber := make(map[int]bool)
			uniqCardNumberList := []int{}
			for _, v := range cards[k] {
				if !uniqCardNumber[v] {
					uniqCardNumber[v] = true
					uniqCardNumberList = append(uniqCardNumberList, v)
				}
			}

			sort.Ints(uniqCardNumberList)

			cards[k] = uniqCardNumberList

			for _, v := range uniqCardNumberList {
				fmt.Fprintln(print, v)
			}
		}
	}
}
