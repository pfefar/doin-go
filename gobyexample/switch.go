package main

import "fmt"

func main() {

	dayofweek := 3

	switch dayofweek {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thurday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}

	hero := "thor"

	switch hero {
	case "spiderman":
		fmt.Println("DC")
	case "ironman", "thor":
		fmt.Println("Marvel")
	default:
		fmt.Println("something")
	}

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("it's a bool")
		case int:
			fmt.Println("it's a integer")
		case string:
			fmt.Println("it's a string")
		default:
			fmt.Printf("waat it's a %T", t)
		}

	}

	whoAmI(true)
	whoAmI("hello")
	whoAmI(2)
	whoAmI(nil)

}
