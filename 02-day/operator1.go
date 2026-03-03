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
	var a int = 5
	fmt.Println(a)
}
