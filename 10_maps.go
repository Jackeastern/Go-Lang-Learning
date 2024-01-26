// https://gobyexample.com/maps
package main

import (
	"fmt"
	"maps" // Maps module
)

func main() {

	m := make(map[string]int) // implementation of a string:int map.
	// make(map[key-type]val-type)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"] // Getting value from map. Kind of reminds me of Python dict.
	fmt.Println("v1:", v1)

	v3 := m["k3"] // If no key then zero value returned. Got it.
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m)) // Len method working on map

	delete(m, "k2") // Removes value
	fmt.Println("map:", m)

	clear(m) // Clear removes all values.
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) { // Awesome functions in maps.
		fmt.Println("n == n2")
	}

	keys := make([]string, 0, len(n))
	for k := range n {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	// vs. Did not work for me. Now keys module in my maps.s
	// keys_2 := maps.Keys(n2)i
	// fmt.Println(keys_2)

}
