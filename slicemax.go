package main

import "fmt"

func main() {
	var b, temp, largerNumber int
	var a []int
	for i := 0; i < 5; i++ {
		fmt.Scan(&b)
		a = append(a, b)
	}
	for _, element := range a {
		if element > temp {
			temp = element
			largerNumber = temp
		}
	}

	fmt.Println(largerNumber)
}
