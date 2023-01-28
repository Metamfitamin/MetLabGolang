package main

import "fmt"

func main() {
	a, b := sumInt(1, 0)
	fmt.Println(a, b)
}

func sumInt(numbers ...int) (length int, sum int) {
	length = len(numbers)
	sum = 0
	for _, number := range numbers {
		sum += number
	}

	return length, sum
}
