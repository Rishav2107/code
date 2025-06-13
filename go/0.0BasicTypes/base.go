package main

import "fmt"

//Go compiles directly to native code — the instructions that your processor (CPU) can execute directly.

func main() {

	//There are two ways to declare a variable
	//1 using var keyword
	var a int

	var (
		b = 2
		f = 2.01
	)

	//2. only within a function
	//short declaration operator
	c := 2

	fmt.Printf("a: %T %[1]v\n", a)
	fmt.Printf("b: %T %[1]v\n", b)
	fmt.Printf("c: %T %[1]v\n", c)
	fmt.Printf("f: %T %[1]v\n", f)

	//a = f //not allowed
	a = int(f)
	fmt.Printf("a: %T %[1]v\n", a)

	//other basic types

	//bool
	//these values are not convertibel to/from integers
	//isValid := 1
	//var isValidSet bool = isValid

	//error
	//a special type with one function, Error()
	//an error may be nil or non-nil

	//pointers
	//are physical address
	//a pointer may be nil or non-nil

	//Intialization
	//Go intializes all variables to zero by default if you don't
	//all numerical types get 0
	//bool gets false
	//string gets ""
	//Everything else get nil:
	//pointers,slices,maps,channels,functions,interfaces

	//Constants
	//only numbers booleans and strings can be constants(immutable)
	const Greeting = "Hello, world!"
	// Greeting cannot be reassigned afterwards.
	var greeting = "Hello, world!"
	// greeting can be reassigned.
	greeting = "Hello, Go!"
	fmt.Print(greeting)
	//All strings in Go are immutable by design.
	//const simply prevents you from reassigning the variable.
	s := "Hello"
	// This is fine (creates a new string)
	s = s + ", world!"
	// This is invalid (can't modify in place)
	// s[0] = 'h'
}

/*
a: int 0
b: int 2
c: int 2
f: float64 2.01
a: int 2
Hello, Go!
*/
