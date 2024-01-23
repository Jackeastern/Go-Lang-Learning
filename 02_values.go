package main

import "fmt"

// getting eeror on main. Had to change to main_2 for some reason. Im guessing go is a global languege. // No
// No main must be used only once in the main program file? // will seee late.
func main() {

	fmt.Println("go" + "lang") // Strings can be concatenated

	fmt.Println("1+1 =", 1+1)         // Arithmatics can be done in line.
	fmt.Println("7.0/3.0 =", 7.0/3.0) // Floats and not the .3333

	fmt.Println(true && false) /// Comparison with bool (AND)
	fmt.Println(true || false) // Comparisons with bool (OR)
	fmt.Println(!true)         // Comaprisons with Bool (Not)
	// main() // running main here to test. Nope main did not run im guessing it may need to be impored.
}
