package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	h := n / 3600
	m := (n % 3600) / 60
	fmt.Print("It is ", h, " hours ", m, " minutes.")
}
