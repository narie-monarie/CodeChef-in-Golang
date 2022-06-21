package main

import (
	"fmt"
)

func solve() {
	var s string
	var input int
	fmt.Scanln(&input)
	fmt.Scanln(&s)
	for i := 0; i < input; i += 2 {
		if s[i] == '0' && s[i+1] == '0' {
			fmt.Printf("A")
		} else if s[i] == '0' && s[i+1] == '1' {
			fmt.Printf("T")
		} else if s[i] == '1' && s[i+1] == '0' {
			fmt.Printf("C")
		} else {
			fmt.Printf("G")
		}
	}

	println()
}
func main() {
	var t int
	fmt.Scanln(&t)
	for i := t; i > 0; i -= 1 {
		solve()
	}
}
