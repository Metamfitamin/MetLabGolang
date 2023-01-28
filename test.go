package main

import "fmt"

func main() {
	var name string
	var a, b, c int
	fmt.Print("Введите имя: ")
	fmt.Scan(&name)
	fmt.Print("Введите возраст: ")
	fmt.Scan(&a, &b, &c)

	fmt.Println(name, a, b, c)
}
