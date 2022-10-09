package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//Standard Library Of Go Language//

	greeting := "Hello There Friends"
	fmt.Println(strings.Contains((greeting), "Hello")) //returns bool value

	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi")) //replaces all the occurences of Hello with Hi

	fmt.Println(strings.ToUpper(greeting)) //converts the string to upper case

	fmt.Println(strings.Index(greeting, "T")) //returns the index of the first occurence of T

	fmt.Println(strings.Split(greeting, " ")) //splits the string into array of strings

	fmt.Println(strings.Count(greeting, "e")) //returns the number of occurences of e

	fmt.Println(strings.Repeat(greeting, 3)) //repeats the string 3 times

	fmt.Println(strings.Trim(greeting, "Hello")) //trims the string from both sides

	fmt.Println(strings.TrimLeft(greeting, "Hello")) //trims the string from left side

	fmt.Println(strings.TrimRight(greeting, "Hello")) //trims the string from right side

	fmt.Println(strings.TrimPrefix(greeting, "Hello")) //trims the string from left side

	fmt.Println(strings.TrimSuffix(greeting, "Hello")) //trims the string from right side

	fmt.Println(strings.Fields(greeting)) //splits the string into array of strings

	fmt.Println(strings.Title(greeting)) //converts the first letter of each word to upper case

	fmt.Println(strings.ToLower(greeting)) //converts the string to lower case

	fmt.Println(strings.HasSuffix(greeting, "Hello")) //returns bool value

	fmt.Println(strings.HasSuffix(greeting, "Hello")) //returns bool value

	fmt.Println(strings.HasPrefix(greeting, "Hello")) //returns bool value

	fmt.Println(strings.HasPrefix(greeting, "Hello")) //returns bool value

	fmt.Println(strings.EqualFold(greeting, "Hello")) //returns bool value

	fmt.Println(strings.Compare(greeting, "Hello")) //returns bool value

	fmt.Println(strings.Index(greeting, "Hello"))

	ages := []int{20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100} //slice of 17 integers

	sort.Ints(ages) //sorts the slice in ascending order

	fmt.Println(ages)

	index := sort.SearchInts(ages, 30) //returns the index of the element 30

	fmt.Println(index)

	names := []string{"rajesh", "kumar", "tapadia", "rajesh"} //slice of 4 strings

	sort.Strings(names) //sorts the slice in ascending order

	fmt.Println(names)

	index = sort.SearchStrings(names, "rajesh") //returns the index of the element rajesh

	fmt.Println(index)

	sort.Reverse(sort.StringSlice(names)) //sorts the slice in descending order

	fmt.Println(names)

}
