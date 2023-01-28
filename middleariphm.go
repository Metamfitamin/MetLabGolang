package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	a = (a + b) / 2
	fmt.Print(a)

}
