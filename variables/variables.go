package variables

import "fmt"

func UseVariables() {
	var msg string
	msg = "Hellow world 1"

	fmt.Println("Variable msg => " + msg)

	msg2 := "Hello 2"

	fmt.Println("Variable msg 2 => " + msg2)
}
