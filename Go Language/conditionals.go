package main

import "fmt"

func main() {
	//Condional statements

	//if
	x := 10
	if x < 5 {
		fmt.Println("x is less than 5")
	}

	//if else

	if x < 5 {
		fmt.Println("x is less than 5")
	} else {
		fmt.Println("x is greater than 5")
	}

	//if else if
	if x<5{
		fmt.Println("x is less than 5")
	} else if x == 5 {
		fmt.Println("x is equal to 5")
	} else {
		fmt.Println("x is greater than 5")
	}

	//switch

	switch x {
		case 1:
			fmt.Println("x is equal to 1")
		case 2:
			fmt.Println("x is equal to 2")
		case 3:
			fmt.Println("x is equal to 3")
		default:
			fmt.Println("x is not equal to 1, 2 or 3")
	

}

}