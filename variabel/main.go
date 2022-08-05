package main

import "fmt"

func main() {
	var firstName string = "john"

	var lastName string
	lastName = "wick"

	fmt.Printf("halo %s %s! \n", firstName, lastName)

	lastNamez := "wick"
	fmt.Printf("halo %s %s! \n", firstName, lastNamez)

	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	fmt.Println(first, second, third)

	name := new(string)
	
	fmt.Println(name)
	fmt.Println(*name)
}