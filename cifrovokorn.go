package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	n = 1 + ((n - 1) % 9)
	fmt.Print(n)

}
