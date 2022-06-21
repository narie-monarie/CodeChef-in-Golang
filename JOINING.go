package main

import "fmt"

func main() {
	var a, n, k int
	fmt.Scanln(&a)
	for i := 0; i < a; i++ {
		fmt.Scanln(&n, &k)
		x := n / 5
		if n%5 != 0 {
			x += 1
		}
		m := k / 5
		if k%5 != 0 {
			m += 1
		}
		fmt.Println(x - m)
	}
}
