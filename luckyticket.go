package main

import "fmt"

func main() {
	var a uint32
	fmt.Scan(&a)
	var b, c uint32 = a / 1000, a % 1000
	b = (b / 100) + (b % 10) + ((b % 100) / 10)
	c = (c / 100) + (c % 10) + ((c % 100) / 10)
	if b == c {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
