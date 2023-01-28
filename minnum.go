package main

import "fmt"

func main() {
	var array []int
	var a, largerNumber, temp, count, n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		array = append(array, a)
		if i == 0 {
			largerNumber = a
		}
		temp = a
		if temp < largerNumber {
			largerNumber = temp
			count = 1
		} else if temp == largerNumber {
			count++
		}
	}
	fmt.Println(count)
}
