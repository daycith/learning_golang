package loops

import "fmt"

func BasicForLoop() {
	for x := 0; x < 10; x++ {
		doSomethingWithX(x)
	}
}

func WhileLoop() {
	x := 0

	for x < 10 {
		doSomethingWithX(x)
		x++
	}
}

func DoWhileLoop() {
	x := 0
	for condition := true; condition; condition = x < 10 {
		doSomethingWithX(x)
		x++
	}
}

func DoWhileLoop2() {
	x := 0
	for {
		doSomethingWithX(x)
		x++

		if x == 10 {
			break
		}
	}
}

func RangeLoop() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, x := range numbers {
		doSomethingWithX(x)
	}
}

func doSomethingWithX(x int) {
	fmt.Printf("The value of x is %v \n", x)
}
