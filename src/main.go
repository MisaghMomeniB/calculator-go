package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: calculator <num1> <operator> <num2>")
		fmt.Println("Operators: + - * /")
		return
	}

	num1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid number:", os.Args[1])
		return
	}

	op := os.Args[2]

	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Invalid number:", os.Args[3])
		return
	}

	var result float64
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: division by zero")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Unsupported operator:", op)
		return
	}

	fmt.Printf("%g %s %g = %g\n", num1, op, num2, result)
}
