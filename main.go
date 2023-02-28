package main

import (
	"fmt"
	"log"

	"github.com/daycith/learning_golang/functions"
	"github.com/daycith/learning_golang/variables"
)

func main() {
	variables.UseVariables()

	sum := functions.Sum(1, 2)

	fmt.Printf("Sum => %v\n ", sum)

	division, err := functions.Divide(2, 0)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Div => %v\n ", division)
}
