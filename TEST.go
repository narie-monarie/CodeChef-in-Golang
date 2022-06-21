package main

import "fmt"

func main() {
	var n int
	for {
		fmt.Scanln(&n)
		if n == 42 {
			break
		}
		fmt.Println(n)
	}
}
