// https://gobyexample.com/variadic-functions
package main

import "fmt"

// Can be called with any number of trailing arguments. Makes me think for *args in python.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	sum(1, 2)    // two arguments.
	sum(1, 2, 3) // 3 args

	nums := []int{1, 2, 3, 4} // Array of 4 items. Also this looks like a slice instead of an array.
	sum(nums...)
	// more_nums := [4]int{2, 3, 4, 5} // Does not work.
	more_nums := []int{2, 3, 4, 5, 6, 7, 8}
	sum(more_nums...)

}
