package main

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

// func main() {
// 	var a int = 5         // first way
// 	var b bool = false    // first way
// 	var c float64 = 2.345 // first way
// 	variableName := 5     // second way -> type is  inferred
// 	var x, y, z int = 1, 2, 3
// 	var t, h = 4, "Hi"

// 	var (
// 		u int
// 		v int = 1
// 		d     = "Hi"
// 	)

// 	// var 1t = 5 // error Variable name never start with numeric

// 	const PI float32 = 3.14

// 	fmt.Println(a)            // first way
// 	fmt.Println(b)            // first way
// 	fmt.Println(c)            // first way
// 	fmt.Println(variableName) // second way
// 	fmt.Println(x, y, z)      // Multiple variable
// 	fmt.Println(t, h)
// 	fmt.Println(d, u, v)
// 	//fmt.Println(k) // error
// 	fmt.Println(PI)
// }

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

// func main() {
// 	var i, j = "Hello", "World"
// 	fmt.Print(i, "\n")
// 	fmt.Print(j, "\n")
// }

// ---------- Basic Types ---------------------------------------
/*
1. Boolean bool -> faqat true yoki false qiymat
2. Satr(String) string -> "" orqali beriladigan qiymatlar
3. Sonli turlar:
	3.1 Butun sonli turlar -> int,int8,int16,int32,int64
	3.2 Haqiqiy sonli turlar float,float8,float16,float32,float64
	3,3 Complex sonli turlar conplex64 complex128
-----------------------------------------------------------------
*/
// First Basix Types for example
func main() {
	// var isActive bool = true // yokida isActive := true
	// //isActive := false
	// fmt.Println(isActive)

	// username := "Abdusamad"
	// fmt.Println(strings.ToUpper(username))

	// var a int8 = 127 // -128 <-> 127  | 8bits/1byte

	// //var a int8 = 127456 // cannot use 127456 (untyped int constant) as int8 value in variable declaration

	// var b int16 = 32767 // -32768 <-> 32767 // 16bits/2byte

	// var c int32 = -2147483648 // -2147483648 <-> 2147483647  // 32bits/4bytes

	// var d int64 = 9223372036854775807 // -9223372036854775808 <-> 9223372036854775807

	// fmt.Println(a)
	// fmt.Println(b)

	// fmt.Println(c)
	// fmt.Println(d)

	agrigate()
}
