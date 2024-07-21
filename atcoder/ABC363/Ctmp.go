package main

import (
	"fmt"
)

// permute は与えられた文字列の全ての並び替えを生成します
func permute(input string) []string {
	var helper func([]rune, int)
	var res []string

	helper = func(arr []rune, n int) {
		if n == 1 {
			res = append(res, string(arr))
			return
		}
		for i := 0; i < n; i++ {
			helper(arr, n-1)
			if n%2 == 1 {
				arr[0], arr[n-1] = arr[n-1], arr[0]
			} else {
				arr[i], arr[n-1] = arr[n-1], arr[i]
			}
		}
	}

	r := []rune(input)
	helper(r, len(r))
	return res
}

func main() {
	str := "abc"
	permutations := permute(str)
	for _, p := range permutations {
		fmt.Println(p)
	}
}
