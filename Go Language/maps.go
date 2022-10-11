package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":  4.99,
		"pie":   7.99,
		"salad": 6.99,
	}

	fmt.Println(menu)
	fmt.Printf("Price of pie  %v\n", menu["pie"])

	//Looping in Map

	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

	//ints as key type

	phonebook := map[int]string{
		123456789: "John",
		987654321: "Doe",
		999999999: "Jane",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[123456789])

	phonebook[123456789] = "John Doe"

	fmt.Println(phonebook)

	//delete from map

	delete(phonebook, 123456789)

	fmt.Println(phonebook)

	//Maps are used to store key value pairs and are unordered and unindexed and can be used to store any type of data and Maps must have same type of keys and values
}
