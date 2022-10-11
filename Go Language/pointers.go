package main

import (
	"fmt"
)

func UpdateNamei(x *string) {
	*x = "wedge"
}

func main() {
	name := "tiffa"
	// name = UpdateNamei(name)

	fmt.Println("Mmeory Address of name is ", &name)

	m := &name // m is a pointer to name//

	fmt.Println("Mmeory Address of m is ", m)

	//derefencing of pointers

	fmt.Println("Value of m is ", *m)

	//changing the value of name using pointer

	UpdateNamei(m)

	fmt.Println(name)

}
