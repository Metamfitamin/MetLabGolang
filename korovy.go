package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	switch {
	case n%10 == 1 && n%100 != 11:
		fmt.Print(n, " korova")
	case n%10 >= 2 && n%10 <= 4 && n%100 != 12 && n%100 != 13 && n%100 != 14:
		fmt.Print(n, " korovy")
	case n%10 == 0 || (n%10 >= 5 && n%10 <= 9) || (n%100 >= 11 && n%100 <= 14):
		fmt.Print(n, " korov")
	}
}
