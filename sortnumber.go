package main

import "fmt"

func main() {
	var a float64
	fmt.Scanf("%f", &a)
	if a <= 0 {
		fmt.Print("число ")
		fmt.Printf("%4.2f", a)
		fmt.Print(" не подходит")
	} else if a > 10000 {
		fmt.Printf("%e", a)
	} else {
		a = a * a
		fmt.Printf("%.4f", a)
	}
}
