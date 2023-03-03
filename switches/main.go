package main

import "fmt"

func main() {
	weekDay := 5

	switch weekDay {
	case 1:
		fmt.Println("Is Monday")
	case 2:
		fmt.Println("Is Tuesday")
	case 3:
		fmt.Println("Is Wednesday")
	case 4:
		fmt.Println("Is Thursday")
	case 5:
		fmt.Println("Is Friday")
		fallthrough
	case 6, 7:
		fmt.Println("Is weekend")

	default:
		fmt.Println("Invalid day")

	}
}
