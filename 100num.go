package main

import "fmt"

func main() {
	var n int
	for fmt.Scan(&n); ; fmt.Scan(&n) { //счет бесконечный
		switch {
		case n < 10:
			continue
		case n > 100:
			break
		}
		fmt.Println(n)
	}
}
