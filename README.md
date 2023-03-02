# Learning golang

#### _This repo contains some exercises about basic topics of golang._

1. [Variables](#Varaiables)
2. [Functions](#Functions)
3. [Loops](#Loops)
4. [Conditionals](#Conditionals)

### Variables

In Go we can declare variables using the long way:

```go
var msg string
msg = "Hellow world 1"
```

In this case, first we declare the variable "msg" with type string. Then we assign the text "Hello world 1" to the msg variable.

Another way is to use the operator :=, which allows to declare and initialize a variable in one line:

```go
msg2 := "Hello world 2"
```

If we don't use both variables, the compilator will throw an error (Yes, Go is against dead code!), so let's simply print them on the console

```go
fmt.Printf("Variable msg => %v\n", msg)
fmt.Printf("Variable msg 2 => %v\n ", msg2)
```

The result on the console will be:

```go
Variable msg => Hellow world 1
Variable msg2 => Hello world 2
```

### Functions

We use the keyword "func" in order to declare a function, followed by its name, arguments and the return type.

```go
func Sum(arg1 int, arg2 int) int {
	return arg1 + arg2
}
```

Notice that the name of the function starts with a capital letter, this is done so because we want the function to be called from other packages. See [Exported names](https://go.dev/tour/basics/3) for more details.

The Sum function has 2 arguments of the same type, so we could declare them specifying the type once:

```go
func Sum(arg1, arg2 int) int {
	return arg1 + arg2
}
```

A function can return more than one type, including an error. Let's create a new function for dividing 2 numbers. As you know, we can not divide by zero, so we want to control that situation and to return an error:

```go
func Divide(arg1, arg2 int) (int,error) {
	if arg2 == 0 {
		return errors.New("can not divide by 0"), 0
	}
	sum := arg1 + arg2
	return nil, sum
}
```

Returning and handling errors in Go is a little different from other languages, so it deserves a new discussion. However, I will show you how to call this simple function and handle the error in a simple way:

```go
division, err := functions.Divide(2, 0)

if err != nil {
	log.Fatal(err.Error())
}

fmt.Printf("Div => %v\n ", division)
```

As you can see, the onus is entirely on the caller to handle the error appropriately.

### Loops

_Go only gives us the traditional "for" loop_, which works like in Java and other programming languages. It uses three components:

- The initialization, x := 0
- The ending condition, x < 10
- The iteration, x++

Example:

```go
for x := 0; x < 10; x++ {
	doSomethingWithX(x)
}
```

In Go there are no while or do-while loops. However we can use the "for" loop instead of them.

#### While

We emulate a while loop when we skip the initialization and iteration sentences:

```go
x := 0
for x < 10{
	doSomethingWithX(x)
	x++
}
```

#### Do while

We can emulate a do-while loop with one these 2 ways:

```go
x := 0
for condition := true; condition; condition = x < 10{
	doSomethingWithX(x)
	x++
}
```

```go
x := 0
for {
	doSomethingWithX(x)
	x++

	if x == 10 {
		break
	}
}
```

Another common use of for loop, is for looping over elements in arrays, maps and slices using __range__

```go
numbers := []int{0,1,2,3,4,5,6,7,8,9}
for _, x := range numbers{
	doSomethingWithX(x)
}
```

If you did not know, we use _ to ignore a returned value from a function (like here), and to import libraries that we will not use directly, but our code needs them implicitly.

#### Conditionals

The if/else conditional works in go as in other languages. Maybe the most important thing to mention is that is nood mandatory to use brackets in the if condition:

```go
myNumber := 15
if myNumber % 2 == 0{
	fmt.Printf("The number %d is even\n", myNumber)
}else{
	fmt.Printf("The number %d is odd\n", myNumber)
}
```