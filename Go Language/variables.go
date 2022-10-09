package main

import "fmt"

//fmt is a package that implements formatted I/O with functions analogous to C's printf and scanf.

func main() {
	//var nameOne string = "John"
	//var nameTwo string = "Doe"
	nameOne := "John" //short hand declaration

	nameTwo := "Doe" //short hand declaration
	var ageOne int = 20
	// var nameThree string

	fmt.Println("Hello, my name is", nameOne, nameTwo, "and I am", ageOne, "years old.")
	fmt.Println("Hello, World!")
	// fmt.Println(nameOne, nameTwo, ageOne, nameThree)

	// fmt.Scanf("%s", &nameThree)

	// fmt.Println("Hi Handsome " + nameThree)

	//datatypes

	//string
	var name string = "John Doe"
	fmt.Println(name)

	//int
	var age int = 20
	fmt.Println(age)

	//float
	var height float32 = 5.11
	fmt.Println(height)

	//bool
	var isCool bool = true
	fmt.Println(isCool)

	//shorthand
	//name := "John Doe"
	//age := 20
	//height := 5.11
	//isCool := true

	var nameshwarOne string = "John Doe"
	var ageOnehwar int = 20
	var heightOnehwar float64 = 5.11
	var isCoolOnehwar bool = true
	nameshwartwo := "kiran chitlapudi"

	fmt.Println(nameshwarOne, ageOnehwar, heightOnehwar, isCoolOnehwar, nameshwartwo)

	//bits and memory

	//int8 //int8 values ranges from -128 to 127
	//int16 //int16 values ranges from -32768 to 32767

	//int32 // int32 values ranges from -2147483648 to 2147483647
	//int64 // int64 values ranges from -9223372036854775808 to 9223372036854775807

	//uint8 // uint8 values ranges from 0 to 255
	//uint16 // uint16 values ranges from 0 to 65535
	//uint32 // uint32 values ranges from 0 to 4294967295
	//uint64 // uint64 values ranges from 0 to 18446744073709551615

	//float32 // float32 values ranges from 1.18e-38 to 3.4e+38
	//float64 // float64 values ranges from 2.23e-308 to 1.8e+308
	//

	//complex64 // complex64 values ranges from (3.4e+38+3.4e+38i) to (3.4e+38+3.4e+38i)
	//complex128 // complex128 values ranges from (1.8e+308+1.8e+308i) to (1.8e+308+1.8e+308i)

	var bakara int8 = 127

	fmt.Println(bakara)

	//constants

	const pi float32 = 3.14
	rajcharan := 3.14
	fmt.Println(pi)
	fmt.Println(rajcharan)

	//pi = 1.2 //error

	//constants with multiple values

	const (
		piOne   = 3.14
		piTwo   = 3.14
		piThree = 3.14
	)

	fmt.Println(piOne, piTwo, piThree)

}
