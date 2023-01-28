package main

import "fmt"

func main() {

	var a int
	fmt.Scan(&a)
	b := a % 3
	a = a / 30
	fmt.Println("It is", a, "hours", b, "minutes.")
}
