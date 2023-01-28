package main

import "fmt"

func main() {

	var first uint64
	var second uint64
	var sum uint64
	second = 1
	var num uint64
	var i uint64
	fmt.Scan(&num)
	for i = 0; ; i++ {
		if first == num {
			fmt.Print(i)
			break
		} else if first > num {
			fmt.Print(-1)
			break
		}
		sum = first + second
		first = second
		second = sum
	}
}
