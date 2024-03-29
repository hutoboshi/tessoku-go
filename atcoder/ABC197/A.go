package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Print(string(s[1]))
	fmt.Print(string(s[2]))
	fmt.Print(string(s[0]))

	fmt.Println()
}
