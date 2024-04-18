package main

import "fmt"

func main() {
	var n, k, r, s, p int
	fmt.Scan(&n, &k, &r, &s, &p)
	var t string
	fmt.Scan(&t)
	hands := make([]rune, 0, len(t))
	var ans int

	for i := 0; i < n; i++ {
		if t[i] == 'r' {
			if i < k {
				ans += p
				hands = append(hands, 'p')
			} else if k <= i && hands[i-k] != 'p' {
				ans += p
				hands = append(hands, 'p')
			} else {
				hands = append(hands, 'x')
			}
		}
		if t[i] == 's' {
			if i < k {
				ans += r
				hands = append(hands, 'r')
			} else if k <= i && hands[i-k] != 'r' {
				ans += r
				hands = append(hands, 'r')
			} else {
				hands = append(hands, 'x')
			}
		}
		if t[i] == 'p' {
			if i < k {
				ans += s
				hands = append(hands, 's')
			} else if k <= i && hands[i-k] != 's' {
				ans += s
				hands = append(hands, 's')
			} else {
				hands = append(hands, 'x')
			}
		}
	}
	fmt.Println(ans)
}
