package main

import "strings"

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"

}

func main() {
	fn, ln := getInitials("gangisetty raj charan")
	println(fn, ln)
	fn2, sn2 := getInitials("gorrilla")
	println(fn2, sn2)
}
