package main

import "fmt"

func main() {
	var n, r int64
	fmt.Scan(&n, &r)

	const M = 1000000007

	//手順1 分詞aを求める
	a := int64(1)
	for i := int64(1); i <= n; i++ {
		a = (a * i) % M
	}

	//手順2 分母bを求める
	b := int64(1)
	for i := int64(1); i <= r; i++ {
		b = (b * i) % M
	}
	for i := int64(1); i <= n-r; i++ {
		b = (b * i) % M
	}

	//手順3 答えを求める
	fmt.Println(division(a, b, M))
}

// aのb乗をmで割った余りを返す関数
func power(a, b, m int64) int64 {
	p, ans := a, int64(1)
	for i := int64(0); i < 30; i++ {
		wari := int64(1 << i)
		if (b/wari)%2 == 1 {
			ans = (ans * p) % m
		}
		p = (p * p) % m
	}
	return ans
}

// a/bをmで割った余りを返す関数
func division(a, b, m int64) int64 {
	return (a * power(b, m-2, m)) % m
}
