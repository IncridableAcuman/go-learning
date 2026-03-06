package main

import (
	"fmt"
)

func arifmetic() {

	var a, b int64
	fmt.Scan(&a, &b)

	addition := a + b // Addition
	var subtraction int64 = 0
	if a > b {
		subtraction = a - b // Substraction
	} else {
		subtraction = b - a // Subtraction
	}

	multiplication := a * b // Multiplication

	division := a / b // Division

	fmt.Println("Addition:", addition)
	fmt.Println("Substaction:", subtraction)
	fmt.Println("Multiplication:", multiplication)
	fmt.Println("Division:", division)
}
