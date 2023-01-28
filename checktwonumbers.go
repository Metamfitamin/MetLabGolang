package main

import (
	"fmt"
	"math"
)

func main() {
	var x, p int
	fmt.Scan(&x)
	fmt.Scan(&p)
	x = reverse(x)
	for x != 0 { //счет бесконечный
		buffer1 := x % 10
		x = x / 10
		//fmt.Println("buffer1", buffer1)
		buffer2 := p
		for buffer2 != 0 {
			buffer3 := buffer2 % 10
			buffer2 = buffer2 / 10
			if buffer3 == buffer1 {
				fmt.Print(buffer3, " ")
			}
		}
		//fmt.Println("x", x)
	}
}
func reverse(x int) int {
	sign := "positive"
	if x >= 0 {
		sign = "positive"
	} else {
		sign = "negative"
	}

	x = int(math.Abs(float64(x)))

	var reversedDigit int

	for x > 0 {
		lastDigit := x % 10
		reversedDigit = reversedDigit*10 + lastDigit

		x = x / 10
	}

	if sign == "negative" {
		reversedDigit = reversedDigit * -1
	}
	return reversedDigit
}
