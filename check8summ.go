package main

import "fmt"

func main() {
	var n, a int
	fmt.Scan(&n)
	summ := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a%8 == 0 && a/100 == 0 && a/10 > 0 {
			summ = summ + a
		}
	}
	fmt.Println(summ)
}
