package main

import "fmt"

func main() {
	var x, p, y int
	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)
	for i := 1; ; i++ { //счет бесконечный
		x = x + (x*p)/100
		if y < x {
			fmt.Println(i)
			break
		}
	}
}
