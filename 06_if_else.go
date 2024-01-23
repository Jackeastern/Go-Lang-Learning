package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Conditionals.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// How about.
	i := 8
	if i%2 == 0 {
		fmt.Println(strconv.Itoa(i) + " is even.")
	} else {
		fmt.Println(strconv.Itoa(i) + " is odd.")
	} // Good? No for some reason the i did not print.

	// If's don't need elses. GOt it.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// || and && Can be used.
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// statement can precede conditionals. that var will be available in all branches.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
