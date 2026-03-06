package main

import "fmt"

type Person struct {
	name      string
	age       int
	job       string
	expirence int
}

func person() {
	var person Person
	person.name = "Izzatbek"
	person.age = 22
	person.job = "Programmer"
	person.expirence = 2
	fmt.Println("Name:", person.name, "Age:", person.age, "Job:", person.job, "Expirence:", person.expirence)
}
