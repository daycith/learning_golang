package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	sum := Sum(1, 2)

	fmt.Printf("Sum => %v\n ", sum)

	division, err := Divide(2, 1)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Div => %v\n", division)
}

func Sum(arg1, arg2 int) int {
	return arg1 + arg2
}

func Divide(arg1, arg2 int) (int, error) {

	if arg2 == 0 {
		return 0, errors.New("can not divide by 0")
	}
	sum := arg1 + arg2
	return sum, nil
}
