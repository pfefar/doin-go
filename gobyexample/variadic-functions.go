package main

import "fmt"

func main() {
	sum(1, 2)
	sum(1, 4, 5)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, i := range nums {
		total += i
	}

	fmt.Println(total)
}
