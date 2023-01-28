package main

import "fmt"

func main() {
	var a uint32
	fmt.Scan(&a)
	if a%400 == 0 || (a%4 == 0 && a%100 != 0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
