/*Variables are divided into two groups: value types and reference types.

//Value types are stored directly in the variable and are copied when passed to a function. Value types include: bool, int, float, string, array, struct, and pointer.

//Reference types are stored in the variable as a pointer to the actual data. Reference types include: slice, map, and channel functions.

*/

package main

import (
	"fmt"
)

func UpdateName(x string) string {
	x = "wedge"
	return x
}

func UpdateMenu(y map[string]float64) map[string]float64 {
	y["tacos"] = 2.99
	return y
}

func main() {
	name := "tifa"
	name = UpdateName(name)
	fmt.Println((name))

	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
		"salad":     7.95,
		"steak":     14.95,
		"tacos":     5.95,
		"burger":    6.95,
	}

	x := UpdateMenu(menu)
	fmt.Println(x)
}
