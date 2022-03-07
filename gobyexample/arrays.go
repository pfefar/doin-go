package main

import "fmt"

func main() {
	var a [4]int
	fmt.Println("arrr", a)

	a[3] = 100
	fmt.Println("arr", a)
	fmt.Println("length: ", len(a))

	dna := [4]string{"A", "T", "C", "G"}
	fmt.Println(dna)

	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)
}
