package main

import "fmt"

func main() {

	mapping := make(map[int]string)
	mapping[1] = "one"
	mapping[2] = "two"
	mapping[3] = "tres"

	fmt.Println(mapping)
	fmt.Println("size: ", len(mapping))
	fmt.Println("third one is ---", mapping[3])

	delete(mapping, 2)
	fmt.Println("after deletion ", mapping)

	_, prs := mapping[4]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
