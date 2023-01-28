package main

import "fmt"

func main() {
	var n, count int
	fmt.Scan(&n)
	for i := 0; i < 3; i++ {
		count = count + n%10
		n = n / 10
	}
	fmt.Print(count)
}
