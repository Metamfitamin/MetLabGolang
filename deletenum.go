package main

import (
	"fmt"
)

func main() {
	var n, buff, a int
	fmt.Scan(&n)
	fmt.Scan(&a)
	var slic []int
	for n != 0 {
		buff = n % 10
		if buff != a {
			slic = append(slic, buff)
		}
		n = n / 10
	}
	for i, j := 0, len(slic)-1; i < j; i, j = i+1, j-1 {
		slic[i], slic[j] = slic[j], slic[i]
	}
	for i := 0; i < len(slic); i++ {
		fmt.Print(slic[i])
	}

}
