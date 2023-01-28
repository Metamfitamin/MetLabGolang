package main

import "fmt"

func main() {
	var a, b, buffer int
	fmt.Scan(&a, &b)
	for i := b; i >= a; i-- {
		buffer = i % 7
		if buffer == 0 {
			fmt.Print(i)
			break
		} else if i == a {
			fmt.Print("NO")

		}
	}
}
