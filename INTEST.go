package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, k int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Scanln(&n, &k)

	counter := 0

	for i := 0; i < n; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		if v%k == 0 {
			counter += 1

		}

	}
	fmt.Println(counter)

}
