package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	if strings.ToUpper(text[:2]) == text[:2] && strings.Contains(text, ".") == true { // проверяет только ecть точка если проверить последний символ лучше
		fmt.Print("Right")                                                              // text[len(text)-1] == 46
	} else {
		fmt.Print("wrong")
	}
}
// проверяет если строка начинается на заглавную и оканчивается на точку.
