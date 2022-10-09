package main

import "fmt"

func main() {

	// age := 32
	// fmt.Print("Hello, World!\n")

	// fmt.Print("Hello, World!")

	// fmt.Println("Hello, World!", age, "years old")

	// //Printf(formatted strings)

	// fmt.Println("My age is %v years old and my height is %f feet tall", age, 5.11)

	//formatting verbs//

	//%v - value in default format
	//%T - type of the value
	//%t - boolean
	//%d - base 10 integer
	//%b - binary
	//%c - character
	//%x - hexadecimal
	//%f - floating point
	//%e - scientific notation
	//%s - string
	//%p - pointer
	//%q - quoted string displays in quotes
	//%#v - value in a Go syntax
	//%T - type of the value
	//%% - literal percent sign

	name := "rajesh "
	// lname := "kumar tapadia"
	agei := "35"

	// fmt.Println("My name is %q %q and %q years old", name, lname, agei) //dislays ascii value of age//

	// fmt.Println("My name is %s %s and %s years old", name, lname, agei) //dislays ascii value of age//

	// fmt.Printf("Age of me is %T", agei) //dislays

	// display type of variable

	// fmt.Printf("Age of me is %T", agei) //dislays

	//Sprintf(formatted strings) saved Formatted strings

	var str = fmt.Sprintf("My name is %v and my name is %v \n", name, agei)
	fmt.Println(str) //saved string//
}
