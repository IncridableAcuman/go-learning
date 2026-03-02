package main

import (
	"fmt"
)

// Ma'lumotlarni chiqarish
// func main() {
// 	var floatValue = 4.5
// 	fmt.Println("Izzatbek Abdusharipov") // spring type
// 	fmt.Printf("%.2f", floatValue)       // float and double type
// 	fmt.Println()                        // enter

// }
/*
func main() {
	var originalProductBalance = 10000 // mahsulotning asl narxi
	var totalProducts = 1000           // bazadagi mahsulotlar miqdori
	var sellingBalance = 12000         // do'kondagi narxi

	var baseOriginalProductTotalBalance = originalProductBalance * totalProducts

	var customerQuantity = 2 // xaridor sotib olmoqchi bo'lgan mahsulotlar soni
	var customerProductBalance = 1

	if customerQuantity > 0 { // Agar customer mahsulot sotib olmoqchi bo'lsa uning qiymatini hisoblab beramiz va bazadagi masulotning miqdoridan customer miqdorini ayiramiz
		customerProductBalance = sellingBalance * customerQuantity // customerning mahsulotini miqdori
		totalProducts -= customerQuantity                          // bazadan ayiramiz

	} else {
		fmt.Println("Mahsulot sotib olinmadi!") // Agar customer mahsulot sotib olmasa beriladigan xabar
	}
	fmt.Println("Mahsulotlarning asl bahosi bilan jami narxi:", baseOriginalProductTotalBalance) // Asl bahosi
	fmt.Println("Bazadigi mahsulotlarning jami soni:", totalProducts)                            // jami soni
	fmt.Println("Xaridorning mahsulotlarini narxi:", customerProductBalance)                     // xaridorning narxi
}
*/
// for sikl
// func main() {
// 	var count = 10
// 	for i := 1; i < count; i++ {
// 		fmt.Println(i, " ")
// 	}

// }

// Variables

func main() {
	var a int = 5         // first way
	var b bool = false    // first way
	var c float64 = 2.345 // first way
	variableName := 5     // second way -> type is  inferred
	var x, y, z int = 1, 2, 3
	var t, h = 4, "Hi"

	var (
		u int
		v int = 1
		d     = "Hi"
	)

	// var 1t = 5 // error Variable name never start with numeric

	fmt.Println(a)            // first way
	fmt.Println(b)            // first way
	fmt.Println(c)            // first way
	fmt.Println(variableName) // second way
	fmt.Println(x, y, z)      // Multiple variable
	fmt.Println(t, h)
	fmt.Println(d, u, v)
	//fmt.Println(k) // error
}

// func main() {
// 	var letter = "Salom Dunyo"
// 	var word = "Hello World"
// 	var halfLetter = "Salom "
// 	var halfWord = "Hello"

// 	fmt.Println(strings.Compare(letter, "Hello World")) // 1 -> false
// 	fmt.Println(strings.Compare(word, "Hello World"))   // 0 -> true

// 	fmt.Println(strings.Contains(letter, halfLetter)) // true
// 	fmt.Println(strings.Contains(word, halfWord))     // true

// 	fmt.Println(strings.Count(letter, halfLetter)) // 1 marta

// 	// fmt.Println(strings.Cut()) //
// }
