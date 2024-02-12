// https://gobyexample.com/multiple-return-values
package main

import "fmt"

// int int shows returns two ints. Got it.
func vals() (int, int) {
	return 3, 7
}

// Here I write my own function
func vals2() (int, string) {
	return 4, "Five"
}

func main() {

	a, b := vals() // Use two vars .
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals() // classic _ blank identifiers.
	fmt.Println(c)

	// Here I add my own.
	z, y := vals2()
	fmt.Println(z)
	fmt.Println(y)
}
