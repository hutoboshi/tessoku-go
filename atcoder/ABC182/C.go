/*
余り
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Split(bufio.ScanWords)
	n := inputInt()

	//nを3で割り切ったあまりを計算
	amariAll := n % 3

	//nを文字列にして変換して１桁ずつチェックする
	nstr := strconv.Itoa(n)
	amari1 := false
	amari2 := false
	for i := 0; i < len(nstr); i++ {
		ketaNum, _ := strconv.Atoi(string(nstr[i]))
		if ketaNum%3 == 1 {
			amari1 = true
		} else if ketaNum%3 == 2 {
			amari2 = true
		}
	}

	if amariAll == 0 {
		fmt.Println(0)
	} else if amariAll == 1 {
		if amari1 {
			if len(nstr) <= 1 {
				fmt.Println(-1)
			} else {
				fmt.Println(1)
			}
		} else {
			if len(nstr) <= 2 {
				fmt.Println(-1)
			} else {
				fmt.Println(2)
			}
		}
	} else if amariAll == 2 {
		if amari2 {
			if len(nstr) <= 1 {
				fmt.Println(-1)
			} else {
				fmt.Println(1)
			}
		} else {
			if len(nstr) <= 2 {
				fmt.Println(-1)
			} else {
				fmt.Println(2)
			}
		}
	}
}

func inputInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func inputString() string {
	scanner.Scan()
	text := scanner.Text()
	return text
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
