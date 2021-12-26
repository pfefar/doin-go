package main

import "fmt"

func main() {
	names := []string{"bob", "alice", "peter"}

	for _, name := range names {
		fmt.Println(name)
	}

	for idx, name := range names {
		fmt.Printf("%s at %d \n", name, idx)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kvs {
		fmt.Printf("%s ---- %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key: ", k)
	}

	// range on strings iterates over Unicode code points.
	// The first value is the starting byte index of the rune and the second the rune itself.
	for i, c := range "pot" {
		fmt.Println(i, c)
	}
}
