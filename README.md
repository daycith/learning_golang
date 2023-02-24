# Learning golang

#### _This repo contains some exercises about basic topics of golang._

1. [Variables](#Varaiables)
2. [Functions](#Functions)

### Variables

In go we can declare variables using the long way:

```go
var msg string
msg = "Hellow world 1"
```
In this case, first we declare the variable "msg" with type string. Then we assign the text "Hello world 1" to the msg variable.

Another way is to use the operator :=, which allows to declare and initialize a variable in one line:

```go
msg2 := "Hello world 2"
```

If we don't use both variables, the compilator will throw an error (Yes, go is against dead code!), so lets simply print them on the console

```go
fmt.Printf("Variable msg => %v", msg)
fmt.Printf("Variable msg 2 => %v", msg2)
```
The result on the console will be:

```sh
Variable msg => Hellow world 1
Variable msg 2 => Hello world 2
````

### Functions

## License

MIT
