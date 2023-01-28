package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var b, c, d int = a % 10, a / 100, (a % 100) / 10
	if b == c || c == d || d == b {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
