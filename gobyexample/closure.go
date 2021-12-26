package main

import "fmt"

func main() {
	nextMul := intSeq()
	fmt.Println(nextMul())
	fmt.Println(nextMul())
	fmt.Println(nextMul())

	newMul := intSeq()
	fmt.Println(newMul())

}

func intSeq() func() int {
	i := 1
	return func() int {
		i *= 2
		return i
	}
}
