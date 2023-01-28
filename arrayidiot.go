package main

import "fmt"

func main() {
	var workArray [10]uint8
	var b uint8
	var changearray [6]uint8
	for i := 0; i < 10; i++ {
		fmt.Scan(&b)
		workArray[i] = b
	}
	for i := 0; i < 6; i++ {
		fmt.Scan(&b)
		changearray[i] = b
	}
	for i := 0; i < 6; i = i + 2 {
		workArray[changearray[i]], workArray[changearray[i+1]] = workArray[changearray[i+1]], workArray[changearray[i]]
	}
	for i := 0; i < 10; i++ {
		fmt.Print(workArray[i], " ")
	}
	for i := 0; i < 6; i++ {
		fmt.Print(changearray[i], " ")
	}
	fmt.Print("type ok")
}
