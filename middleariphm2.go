package main

import "fmt"

func main() {
	var a, b, i float64

	for i = 0; ; i++ {
		fmt.Scan(&a)
		b = b + a
		if a == 0 {
			b = b / i
			fmt.Print(b)
			break
		}
	}
}
