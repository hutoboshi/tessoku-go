package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)

	if len(S) < 3 {
		fmt.Println("No")
		return
	}

	if S[len(S)-1] == '>' && S[0] == '<' {
		var cnt int = 0
		for i := 1; i < len(S)-1; i++ {
			if S[i] == '=' {
				cnt++
			}
		}

		if cnt == len(S)-2 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	} else {
		fmt.Println("No")
	}
}
