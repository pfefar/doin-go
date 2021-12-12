package main

import (
	"fmt"
	"math"
)

const s string = "Novi"

func main() {
	const n = 20

	const d = 6e+11 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	// number n is numeric constant no type, it  will be given a type in the function call context because function call expects it.
	fmt.Printf("%T", math.Sin(n))
}
