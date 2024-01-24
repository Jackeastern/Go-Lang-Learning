// https://gobyexample.com/arrays
package main

import "fmt"

func main() {

	var a [5]int // Arrays are zeo valued. [0,0,0,0,0]
	fmt.Println("emp:", a)

	a[4] = 100 // Index start at 0 [0, 0, 0, 100]
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5} // One liner
	fmt.Println("dcl:", b)

	// How about arrays with strings?
	string_array := [4]string{"I", "am", "string", "Array"}
	fmt.Println("str1", string_array)
	// Works.

	var twoD [2][3]int // 2-D array.
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
