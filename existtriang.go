package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a < b+c && b < a+c && c < a+b {
		fmt.Print("Существует")
	} else {
		fmt.Println("Не существует")
	}

}
