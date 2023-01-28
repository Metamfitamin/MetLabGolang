package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	switch {
	case a < 0:
		fmt.Print("Число отрицательное")
	case a == 0:
		fmt.Print("Ноль")
	case a > 0:
		fmt.Print("Число положительное")
	}
}
