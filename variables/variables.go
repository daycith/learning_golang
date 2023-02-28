package variables

import "fmt"

func UseVariables() {
	var msg string
	msg = "Hellow world 1"

	fmt.Printf("Variable msg => %v\n", msg)

	msg2 := "Hello world 2"

	fmt.Printf("Variable msg 2 => %v\n ", msg2)
}
