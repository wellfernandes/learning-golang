package main

import (
	"fmt"
	"learninggo/calc"
)

func main() {

	var num1 int = 4
	num2 := 2

	result, err := calc.Div(num1, num2)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(result)
}
