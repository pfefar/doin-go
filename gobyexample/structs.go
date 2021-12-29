package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42

	return &p

}

func main() {

	fmt.Println(person{"Jamie", 20})
	fmt.Println(person{name: "Alice", age: 22})
	fmt.Println(person{name: "yoy"}) // ommited field will be zero valued
	fmt.Println(&person{name: "yoy", age: 42})
	fmt.Println(newPerson("jon"))

	s := person{name: "enzo", age: 12}
	fmt.Println(s)

	sp := &s
	fmt.Println(sp.name)

	sp.age = 52

	fmt.Println(s)

}
