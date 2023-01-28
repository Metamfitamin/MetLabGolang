package main

import "fmt"

func main() {
	var a, b uint16
	fmt.Scan(&a)
	fmt.Scan(&b)
	var summ uint16 = 0
	for i := a; i <= b; i++ {
		summ += i
	}
	fmt.Println(summ)

}
