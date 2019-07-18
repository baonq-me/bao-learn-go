package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Require 3 arguments !")
		return
	}

	fmt.Println(os.Args)

	var a, _ = strconv.ParseFloat(os.Args[1], 10)
	var b, _ = strconv.ParseFloat(os.Args[3], 10)
	var op = os.Args[2]

	switch op {
	case "+":
		fmt.Println(a, op, b, "=", a+b)
	case "-":
		fmt.Println(a, op, b, "=", a-b)
	case "*":
		fmt.Println(a, op, b, "=", a*b)
	case "/":
		fmt.Println(a, op, b, "=", a/b)
	default:
		fmt.Println("Invalid operator. Allowed operators: + - * /")
	}
}
