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
	for i := 0; i < n; i += 2 {
		fmt.Print(myArray[i], " ")
	}
}
