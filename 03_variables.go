package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2 // Able to declare multiple.
	fmt.Println(b, c)

	var x, y, z = "a", "b", "c" // my test
	fmt.Println(x, y, z)        // does it work? Yes!

	var d = true // Infers types. Got it.
	fmt.Println(d)

	var e int // Zero  valued if not declared with a value. Got it!
	fmt.Println(e)

	// What about Strings?
	//test_check_string := string
	// var test_check_string_two = string
	//fmt.Printf(test_check_string)
	// fmt.Printf(test_check_string_two)
	// Looks like these don't work. :(

	f := "apple" // Shorthand available in functions. Got it!
	fmt.Println(f)
}
