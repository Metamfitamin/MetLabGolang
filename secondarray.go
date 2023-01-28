package main

import "fmt"

func main() {
	var n, a int
	myArray := [100]int{}
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		myArray[i] = a

	}
	num := 0
	for i := 0; i < n; i++ {
		if myArray[i] >= 0 {
			num++
		}
	}
	fmt.Println(num)
}
