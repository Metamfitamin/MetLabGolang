package main

import (
	"fmt"
)

type shooter struct {
	ON          bool
	Ammo, Power int
}

func (testStruct *shooter) Shoot() bool {
	if testStruct.Ammo > 0 && testStruct.ON == true {
		testStruct.Ammo--
		return true
	} else {
		return false
	}
}
func (testStruct *shooter) RideBike() bool {
	if testStruct.Power > 0 && testStruct.ON {
		testStruct.Power--
		return true
	} else {
		return false
	}
}

func main() {
	c := shooter{true, 1, 2}
	testStruct := &c

	fmt.Println(testStruct.Shoot())
	fmt.Println(testStruct.RideBike())
}
