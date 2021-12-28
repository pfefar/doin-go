package main

import "fmt"

func main() {
	changeIt := 100
	fmt.Println(changeIt)

	simplefunc(changeIt)
	fmt.Println(changeIt)

	pointerfunc(&changeIt)
	fmt.Println(changeIt)

	fmt.Println(&changeIt)
}

func simplefunc(val int) {
	val = 0
}

func pointerfunc(val *int) {
	*val = 0
}
