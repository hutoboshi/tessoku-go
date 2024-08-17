package main

import "fmt"

func rangeFive(yield func() bool) {
	if !yield() {
		return
	}
	if !yield() {
		return
	}
	if !yield() {
		return
	}
	if !yield() {
		return
	}
	if !yield() {
		return
	}
}

func main() {
	for range rangeFive {
		fmt.Println("Hello")
	}
}
