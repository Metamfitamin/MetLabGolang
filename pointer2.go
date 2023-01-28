package main

import "fmt"

func changer(i *int, i1 *int) {
	*i = *i * *i1
}
func main() {
	i := 11
	i1 := 22
	changer(&i, &i1)
	fmt.Println(i)
}
