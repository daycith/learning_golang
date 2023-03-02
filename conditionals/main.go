package main

import "fmt"

func main() {
	myNumber := 15
	if myNumber%2 == 0 {
		fmt.Printf("The number %d is even\n", myNumber)
	} else {
		fmt.Printf("The number %d is odd\n", myNumber)
	}
}
