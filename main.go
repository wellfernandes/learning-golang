package main

import (
	"fmt"
	"learninggo/calc"
)

func main() {

	var num1 int = 4
	num2 := 2

	result := calc.Div(num1, num2)

	fmt.Println(result)
}
