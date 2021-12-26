package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("the number is even")
	} else {
		fmt.Println("number is odd")
	}

	if 8%4 == 0 {
		fmt.Println("the number is divisible by 4")
	}

	// A statement can precede conditionals; any variables declared in this statement are available in all branches.
	if some := 7; some%2 == 0 {
		fmt.Println("woah")
	} else if some%3 == 0 {
		fmt.Println("three woo")
	} else {
		fmt.Println("just some number man")
	}
}
