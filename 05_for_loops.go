package main

import "fmt"

func main() {
	i := 1 // Set var outside loop.
	// Makes this act like a while loop. Conditional loop.
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Traditional loop. Know this from Java.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// Nice for loop function just on the break statement. That's new to me.
	for {
		fmt.Println("loop")
		break
	}

	// Yes the good old continue. Just makes the loop skip the rest. I've got a good
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
