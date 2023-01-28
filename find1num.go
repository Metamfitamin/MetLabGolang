package main

import "fmt"

func main() {
	var a uint16
	fmt.Scan(&a)
	/*switch {
	case a/10000 > 0:
		println(a / 10000)
	case a/10000 == 0 && a/1000 > 0:
		fmt.Println(a / 1000)
	case a/10000 == 0 && a/1000 == 0 && a/100 > 0:
		fmt.Println(a / 100)
	case a/10000 == 0 && a/1000 == 0 && a/100 == 0 && a/10 > 0:
		fmt.Println(a / 10)
	case a/10000 == 0 && a/1000 == 0 && a/100 == 0 && a/10 == 0:
		fmt.Println(a % 10)
	}*/ //типа неправильно
	switch {
	case a > 9 && a < 100:
		fmt.Println(a / 10)
	case a > 99 && a < 1000:
		fmt.Println(a / 100)
	case a > 999 && a < 10000:
		fmt.Println(a / 1000)
	case a == 10000:
		fmt.Println(a / 10000)
	case a < 10:
		fmt.Println(a)
	}
}
