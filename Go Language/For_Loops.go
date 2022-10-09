package main

import "fmt"

func main() {
	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	names := []string{"rajesh", "kumar", "tapadia", "rajesh"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for i := range names {
		fmt.Println(i, names[i])
	}

	for index, value := range names {
		fmt.Println(index, value)
	}

	for _, value := range names { //to get rid of index
		fmt.Println(value)
		value = "Baby "
	}
}
