package main

import (
	"fmt"
)

func main() {

	/*
		Ma'lumot turlari asosan 4 xil bo'ladi.
			1. Asosiy turlar ya'ni Basic types
			2. Agregat turlari ya'ni Aggregate types
			3. Havola turlari ya'ni Reference types
			4. Interface turlari ya'ni Interface types

	*/
	// 1.  --------------- < BASIC TYPES > -------------------------
	/*
		Asosiy turlar o'z ichiga quyidagilarni oladi.
			1.1. Mantiqiy tur (Boolean) --> bool;
			Faqatgina true yokida false qiymatni qabul qiladi.

			1.2. Sonli turlar (Numeric)
				1.2.1 Butun sonlar --> int,int8,int16,int32,int64 bular ishorali yokida uint,uint8,uint16,uint32,uint64 ishorasiz
				1.2.2 Haqiqiy sonlar --> float64, float64
				1.2.3 Complex sonlar --> complex64,complex128

			1.3. Satrli tur (String)
			var username string = "Izzatbek"

	*/

	// mantiqit
	var isActive bool = true
	fmt.Println(isActive)

	// Butun sonlar
	var a int = 5
	var b int8 = -128
	var c int16 = 255
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Haqiqiy sonlar
	var haqiqiySon float32 = 2345.34
	fmt.Println(haqiqiySon)
	fmt.Printf("Number value: %v and type: %T\n", haqiqiySon, haqiqiySon)

	// Complex son
	c1 := complex(2, 3)
	fmt.Println(c1)

	// Satrli tur

	firstName := "Husan"
	lastName := "Hasanov"

	fmt.Println(firstName, lastName)
	//fmt.Print(firstName, "\n", lastName)

	//fmt.Printf("Firstname has value: %v and type: %T\n", firstName, firstName)

	// Formatting Verbs for Printf()

	integer := 12
	float := 1.2
	name := "Mustang"
	//  Bu asosan qiymatni standart formatda chiqaradi
	fmt.Printf("Integer value: %v\n", integer)
	fmt.Printf("Float value: %v\n", float)
	fmt.Printf("String value: %v\n", name)

	// O'zgaruvchining turi uchun %T dan foydalanamiz

	fmt.Printf("Integer type: %T\n", integer)
	fmt.Printf("String type: %T\n", name)
}
