package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Simple Calculator!")

	var num1, num2 float64
	var operation string
	
	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)
	
	fmt.Print("Enter the operation (+, -, *, /): ")
	fmt.Scan(&operation)

	var result float64
	var err error
	switch operation {
	case "+": result = add(num1, num2)
	case "-": result = subtract(num1, num2)
	case "*": result = multiply(num1, num2)
	case "/": 
		result, err = divide(num1, num2)
		if err!=nil {
			fmt.Println("Error:", err)
			return
		}

	default: 
		fmt.Println("Invalid operation!")
		return
	}

	fmt.Printf("Result: %.2f\n", result)
}


func add(a, b float64) float64 {
	return a+b
}

func subtract(a, b float64) float64 {
	return a-b
}

func multiply(a,b float64)  float64 {
	return a*b
}

func divide(a,b float64) (float64, error) {
	if b==0{
		return 0, errors.New("Division by zero is not allowed")
	}
	return a/b, nil
}


