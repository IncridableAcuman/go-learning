package main

import (
	"fmt"
)

/*
Go dasturlash tilida array bilan ishlash.

Arrays yokida massivs - bu har bir qiymat uchun alohida o'zgaruvchi elon qilmasdan balki bitta o'zgaruvchida
bitta turdagi bir nechta ma'lumotlarni saqlashga aytiladi.
Go dasturlash tilida arraylar ikki xil yo'l bilan elon qilinadi.

------------------------One way---------------------------------------
	var fruits = [3]stirng{"Apple","Banana","Peach"} ---> first
	fruits := [3]{1,2,3}

*/

func main() {
	// ----------- < FIST > --------------------------
	var fruits = [3]string{"Orange", "Banana", "Apple"}
	fmt.Println("Array", fruits)
	fmt.Println("Array of size:", len(fruits))
	fmt.Println(fruits[0]) // first element

	for i := 0; i < len(fruits); i++ {
		fmt.Print(i, "-element:", fruits[i], "\n")
	}

	// -------------- < SECOND WAY > -----------------
	fmt.Println()
	cars := [3]string{"Nexia", "Gentra", "Malibu"}

	for j := 0; j < len(cars); j++ {
		fmt.Print(j, "-element:", cars[j], "\n")
	}

	// ----------- < Find a minimum element from array > ---------------------

	var num int
	fmt.Print("N=")
	fmt.Scan(&num)
	arrays := make([]int64, num)

	for i := 0; i < num; i++ {
		fmt.Scan(&arrays[i])
	}
	minElement := arrays[0]

	for i := 0; i < len(arrays); i++ {
		if arrays[i] < minElement {
			minElement = arrays[i]
		}
	}
	fmt.Println(minElement)
}
