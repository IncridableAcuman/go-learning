package main

import (
	"fmt"
)

func bitwise() {
	x := 3 & 5 // AND
	fmt.Println(x)
	x = 4 | 3 // OR
	fmt.Println(x)
	x = 3 ^ 5 // XOR
	fmt.Println(x)
	x = 3 << 2 // Zero fill left shift
	fmt.Println(x)
	x = 4 >> 2 // Signed right shift
	fmt.Println(x)
}
