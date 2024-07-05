package main

import "fmt"

func main() {
	var s1, s2 string
	fmt.Scan(&s1)
	fmt.Scan(&s2)
	if (s1 == ".#" && s2 == "#.") || (s1 == "#." && s2 == ".#") {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
