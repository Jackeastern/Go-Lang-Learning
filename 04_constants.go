// Learning about constants. Reminds be a lot about JS.
package main

// Looks like you can import multiple packages at once like shown below.
import (
	"fmt"
	"math"
)

const s string = "constant" // Wonder if this is also called a global var?

func main() {
	fmt.Println(s) // Globals available in functions.

	const n = 500000000 // numeric constant has no type until it’s given one. Got it.

	const d = 3e20 / n //  expressions perform arithmetic with arbitrary precision. Assuming that means within the limits of the given hardware and context.
	fmt.Println(d)

	fmt.Println(int64(d)) // types can be given.

	fmt.Println(math.Sin(n)) // Types can be given during function calls.
}
