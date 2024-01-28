// Learn about functions
// https://gobyexample.com/functions

package main

import "fmt"

// Function 1 adds a and b ar returns sum
func plus(a int, b int) int {

	return a + b
}

// Function 2 ads a, b, c and returns sum
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

}
