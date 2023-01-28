package main

import "fmt"

func main() {
	var n int
	summ := 0
	buffer := 0
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) { //счет до нуля
		if buffer < n {
			buffer = n
			summ = 1
		} else if buffer == n {
			summ++
		}
	}
	fmt.Println(summ)
}
