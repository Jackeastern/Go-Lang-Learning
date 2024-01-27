// https://gobyexample.com/range
// Learning about range
package main

import "fmt"

func sum_values(array_of_numbers []int) int { // Had to google how to accept arrays. Let test it out.
	sum := 0
	for _, array_number := range array_of_numbers {
		sum += array_number
	}
	return sum
}

func main() {

	nums := []int{2, 3, 4}     // This would be an array. Right? Correct
	sum := 0                   // Variable to track sum value
	for _, num := range nums { // _ because you don't worry about keeping that pos value. Same as with python.
		sum += num // Going through a loop and adding the value to the sum that is in the range. Would it work if I do a function for that?
	}
	fmt.Println("sum:", sum)

	fmt.Println("Testing my sum function:", sum_values(nums)) // Learning us of own functions 😀

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i) // Keeping the index dumber. Got it. Almost like enumerate.
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"} // Theres that one liner map. Yay.
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v) // can loop through key, value almost like zip in python. Also noting the formatting of inserting variables in strings.
	}

	for k := range kvs {
		fmt.Println("key:", k) // wil only print first value got it. In this case the key.
	}

	// so that means
	for _, v := range kvs {
		fmt.Println("value:", v) // this would only print the values.
	}

	for i, c := range "go" {
		fmt.Println(i, c) // from the output looks like ascii value is printed. Nope unicode.
	}
}
