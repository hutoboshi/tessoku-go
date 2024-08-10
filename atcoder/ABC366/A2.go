package main

import (
	"fmt"
)

func main() {
	var N, T, A int
	fmt.Scan(&N, &T, &A)

	remainingVotes := N - (T + A)

	if T > A+remainingVotes || A > T+remainingVotes {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
