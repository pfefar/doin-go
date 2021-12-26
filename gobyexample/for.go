package main

import "fmt"

func main() {

	i := 1

	for i <= 4 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 1; j < 11; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("before breaking out")
		break
	}

	for n := 0; n < 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
