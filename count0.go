package main

import "fmt"

func main() {
	var n, b, count int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&b)
		if b == 0 {
			count++
		}
	}
	fmt.Print(count)

}
