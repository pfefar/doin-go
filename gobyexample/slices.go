package main

import "fmt"

func main() {
	slicey := make([]int, 3)
	fmt.Println(slicey)

	slicey[0] = 111
	slicey[1] = 222
	slicey[2] = 333
	fmt.Println(len(slicey))
	fmt.Println("get: ", slicey[2])

	newSlicey := append(slicey, 444)
	fmt.Println(newSlicey)

	copySlicey := make([]int, len(newSlicey))
	newSlicey = append(newSlicey, 555)
	fmt.Println("newSlicey: ", newSlicey)
	copy(copySlicey, newSlicey)
	fmt.Println("copy: ", copySlicey)

	l := newSlicey[0:2]
	fmt.Println(l)

	l = newSlicey[:3]
	fmt.Println(l)

	l = newSlicey[3:]
	fmt.Println(l)

	dnaSlice := []string{"A", "T", "C", "G"}
	fmt.Println(dnaSlice)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
