package main

import "fmt"

func main() {
	var ages [3]int = [3]int{20, 25, 30} //array of 3 integers
	fmt.Println(ages)

	var agesi = [3]int{20, 25, 30} //array of 3 integers

	fmt.Println(agesi)

	fmt.Println(ages, agesi, len(ages), len(agesi))

	names := [4]string{"rajesh", "kumar", "tapadia", "rajesh"} //array of 4 strings

	fmt.Println(names, len(names))

	//slices

	var scores = []int{100, 50, 60} //slice of 3 integers

	fmt.Println(scores, len(scores))

	scores = append(scores, 85) //append 85 to slice and updates the dtring

	fmt.Println(scores, len(scores))

	//slice range//

	rangeOne := names[1:3] //slice of names from index 1 to 3

	fmt.Println(rangeOne)

	rangeTwo := names[1:] //slice of names from index 1 to end

	fmt.Println(rangeTwo)

	rangeThree := names[:3] //slice of names from index 0 to 3

	fmt.Println(rangeThree)

	rangeFour := names[:] //slice of names from index 0 to end

	fmt.Println(rangeFour)

}
