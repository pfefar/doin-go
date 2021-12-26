package main

import "fmt"

func main() {

	one, two := vals()
	fmt.Println(one)
	fmt.Println(two)

	_, two = vals()
	fmt.Println(two)
}

func vals() (int, int) {
	return 4, 5
}
