package main

import (
	"fmt"
	"introduction/calc"
)

func MainMain() int {
	fmt.Println("Version 5.0.0")
	return 10000
}

func main() {
	defer fmt.Println("deffer the running of this print...")
	fmt.Println("Main is running...")
	fmt.Println(calc.CalcCustom())

	fmt.Println(calc.Newfunction1())

	fmt.Println(calc.Newfunction2())
}
