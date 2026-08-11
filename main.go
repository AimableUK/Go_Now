package main

import "fmt"

func main() {
	// Hello World
	fmt.Println("Hello, World!")

	// variables
	var stu1 string = "John"
	var stu2 = "Peter"

	// this can't be done outside a function as it can be done for thi: var stu1 string = "John"
	x := 2

	fmt.Print(stu1, "\n")
	fmt.Print(stu2)
	fmt.Println(x)

	// value assignment
	var stud string
	stud = "John"
	fmt.Println(stud)

	// Multivariable
	var a, b, c = 4, "hiii", 5
	var v, y, z int = 4, 5, 6
	d, e := 5, "world!"
	var (
		n int
		m int
		o string = "hellow"
	)

	fmt.Println("multiVariable")
	fmt.Println(a, b, c, v, y, z, d, e, n, m, o)

	// camel case:
	// Each word, except the first, starts with a capital letter:
	// myVariableName = "John"

	// Pascal Case
	// Each word starts with a capital letter:
	// MyVariableName = "John"

	// Snake Case
	// Each word is separated by an underscore character:
	// my_variable_name = "John"

	// CONSTANTS
	// Untyped constant
	const PI = 3.14
	// Typed constant
	const MODEL string = "TOYOTA"
	fmt.Println(PI)

	const (
		A int = 1
		B     = 3.14
		C     = "Hi!"
	)

	// Printf
	var ij string = "Hello"
	var j int = 15

	fmt.Printf("i has the value: %v and type: %T\n", ij, ij)
	fmt.Printf("j has the value: %#v and type: %T\n", j, j)

	var i = 15.5
	var txt = "Hello World!"
	var abc = "hello"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)

	fmt.Printf("%q\n", abc)

}
