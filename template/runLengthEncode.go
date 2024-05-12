package main

import "strconv"

func main() {

}

// 使い方はABC259 Cを参照
func runLengthEncode(s string) (string, [][]int) {
	ret := make([]byte, 0)
	ret2 := make([][]int, 0)
	cnt := 0
	t := []int{0, 0}
	for i := 0; i < len(s); i++ {
		if i == 0 {
			ret = append(ret, s[i])
			t[0] = int(s[i])
			cnt++
			continue
		}
		if s[i] == s[i-1] {
			cnt++
			continue
		}
		ret = append(ret, []byte(strconv.Itoa(cnt))...)
		t[1] = cnt
		ret2 = append(ret2, t)
		t = []int{int(s[i]), 0}
		cnt = 1
		ret = append(ret, s[i])
	}
	ret = append(ret, []byte(strconv.Itoa(cnt))...)
	t[1] = cnt
	ret2 = append(ret2, t)

	return string(ret), ret2
}
