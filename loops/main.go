package main

import "fmt"

func main() {
	basicForLoop()
	fmt.Println("_____________________________________")
	whileLoop()
	fmt.Println("_____________________________________")
	doWhileLoop()
	fmt.Println("_____________________________________")
	doWhileLoop2()
	fmt.Println("_____________________________________")
	rangeLoop()
}

func basicForLoop() {
	for x := 0; x < 10; x++ {
		doSomethingWithX(x)
	}
}

func whileLoop() {
	x := 0

	for x < 10 {
		doSomethingWithX(x)
		x++
	}
}

func doWhileLoop() {
	x := 0
	for condition := true; condition; condition = x < 10 {
		doSomethingWithX(x)
		x++
	}
}

func doWhileLoop2() {
	x := 0
	for {
		doSomethingWithX(x)
		x++

		if x == 10 {
			break
		}
	}
}

func rangeLoop() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, x := range numbers {
		doSomethingWithX(x)
	}
}

func doSomethingWithX(x int) {
	fmt.Printf("The value of x is %v \n", x)
}
