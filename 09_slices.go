// Learning about slices at https://gobyexample.com/slices
// Also read https://go.dev/blog/slices-intro
package main

// Is slices an imported module...?
import (
	"fmt"
	"slices"
)

func main() {

	// An uninitialized slice equals to nil and has length 0.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3) //  empty slice with non-zero length, use the builtin make
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a" // Get an set the sam
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d") // Append like this
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s) // Make a copy
	fmt.Println("cpy:", c)

	l := s[2:5] // Splicing
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"} // One line
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2") // comparing
	}

	twoD := make([][]int, 3) // Multi Dimensional
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // Print looks like Array but is slice.
}
